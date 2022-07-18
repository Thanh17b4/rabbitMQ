package handler

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
)

var jwtKey = []byte("secretKey")

type TokenUser struct {
	userService UserService
}

func NewToken(userService UserService) *TokenUser {
	return &TokenUser{userService: userService}
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type Claims struct {
	Username string
	jwt.StandardClaims
}

func (th *TokenUser) CreatToken(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]
	id, er := strconv.ParseInt(userID, 10, 64)
	if er != nil {
		fmt.Println("userID must be number: ", er.Error())
		return
	}
	user, errs := th.userService.GetDetailUser(id)
	if errs != nil {
		fmt.Println("id is not valid: ", errs.Error())
		return
	}
	user1 := user.Username
	pass1 := user.Password
	credentials := &Credentials{
		Username: user1,
		Password: pass1,
	}
	var users = map[string]string{
		user1:       pass1,
		"user2":     "password2",
		"thanh17b4": "22121992Th",
	}
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		fmt.Println("had an error: ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	expectedPassword, ok := users[credentials.Username]
	if !ok || expectedPassword != credentials.Password {
		fmt.Println("password is not correct: ", err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	CreatAt := time.Now()
	expirationTime := time.Now().Add(time.Minute * 1)
	claims := Claims{
		Username: credentials.Username,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  CreatAt.Unix(),
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  expirationTime,
		Path:     "/",
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	})
	json.NewEncoder(w).Encode(tokenString)
	return
}
func (th *TokenUser) VerifyToken(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			fmt.Println("aa")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tokenStr := cookie.Value
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			fmt.Println("bb")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		fmt.Println("cc")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	w.Write([]byte(fmt.Sprintf("Welcome %s!", claims.Username)))
}

func (th *TokenUser) Refresh(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			fmt.Println("dd")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tknStr := c.Value
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"massage": "time is over",
		})
		return
	}
	expirationTime := time.Now().Add(1 * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(tokenString)
	http.SetCookie(w, &http.Cookie{
		Name:    "refresh_token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}
