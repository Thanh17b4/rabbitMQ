package handle

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

var jwtKey = []byte("secretKey")

//type User struct {
//	userService UserService
//}

//func NewToken(userService UserService) *User {
//	return &User{userService: userService}
//}
// thong tin can xac thuc
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// thong tin yeu cau
type Claims struct {
	Username string
	jwt.StandardClaims
}

func CreatToken(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//userID := vars["id"]
	//id, error := strconv.ParseInt(userID, 10, 64)
	//if error != nil {
	//	fmt.Println("userID must be number: ", error.Error())
	//	return
	//}
	//user, errs := lh.userService.GetDetailUser(id)
	//if errs != nil {
	//	fmt.Println("id is not valid: ", errs.Error())
	//	return
	//}
	//json.NewEncoder(w).Encode(user)
	//fmt.Println("", user)
	var credentials Credentials
	var users = map[string]string{
		//user.Username: user.Password,
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
		fmt.Println("password is not valid")
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
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
	json.NewEncoder(w).Encode(tokenString)
	return
}
func VerifyToken(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tokenStr := c.Value
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
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
	w.Write([]byte(fmt.Sprintf("Welcome %s!", claims.Username)))
}

func Refresh(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
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
	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 90*time.Second {
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
	http.SetCookie(w, &http.Cookie{
		Name:    "refresh_token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}
