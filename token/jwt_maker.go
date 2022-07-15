package token

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

var jwtKey = []byte("secretKey")

type Cred struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type Claims struct {
	Username string
	jwt.StandardClaims
}

func Creat(w http.ResponseWriter, r *http.Request) {
	var User = map[string]string{
		"thanh17b4": "22121992Th",
	}
	var cred Cred
	err := json.NewDecoder(r.Body).Decode(&cred)
	if err != nil {
		fmt.Println("could not decode Cred: ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	expectedPassword, ok := User[cred.Username]
	if !ok || expectedPassword != cred.Password {
		fmt.Println("password is not correct: ", err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	issueAt := time.Now()
	expirationTime := issueAt.Add(time.Minute * 1)
	claims := &Claims{
		Username: cred.Username,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  issueAt.Unix(),
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		fmt.Println("", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}
func Verify(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			fmt.Println("no cookie is found: ", err.Error())
			return
		}
		w.WriteHeader(http.StatusBadRequest)
	}
	token := cookie.Value
	claims := &Claims{}
	tokenString, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		fmt.Println("", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tokenString.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	w.Write([]byte(fmt.Sprintf("wellcome %s ", claims.Username)))
}

func Refresh(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		fmt.Println("could not get cookie: ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	token := cookie.Value
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			fmt.Println("", err.Error())
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		fmt.Println("had an error: ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		fmt.Println("token is not valid: ", err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > time.Second*30 {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"massage": "time is over",
		})
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	issueAt := time.Now()
	expirationTime := issueAt.Add(time.Minute * 1)
	claims.IssuedAt = issueAt.Unix()
	claims.ExpiresAt = expirationTime.Unix()
	NewToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	NewTokenString, err := NewToken.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:  "refresh_token",
		Value: NewTokenString,
	})
}
