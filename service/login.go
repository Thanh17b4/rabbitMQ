package service

import (
	"errors"
	"github.com/Thanh17b4/practice/model"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type LoginService struct {
	otpRepo  OtpRepo
	userRepo UserRepo
}

var jwtKey = []byte("secretKey")

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type Claims struct {
	Username string
	jwt.StandardClaims
}

func NewLogin(otpRepo OtpRepo, userRepo UserRepo) *LoginService {
	return &LoginService{otpRepo: otpRepo, userRepo: userRepo}
}

func (l LoginService) Login(email string, password string) (string, error) {
	user, err := l.userRepo.GetUserByEmail(email)
	if email != user.Email {
		return "", errors.New("could not find email in database")
	}
	errs := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if errs != nil {
		return "", errors.New("password is not correct")
	}
	if user.Activated == 0 {
		return "", errors.New("account has not been activated, please activate your account")
	}
	// generate a token here
	now := time.Now()
	claims := model.Claims{
		Username: user.Username,
		Email:    email,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  now.Unix(),
			ExpiresAt: now.Add(time.Minute * 2).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(model.JwtSecretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (l LoginService) Refresh(token string) (string, error) {
	//var r *http.Request
	//tknStr := r.Header.Get("Authorization")
	//tokenArray := strings.Split(tknStr, " ")
	//if len(tokenArray) != 2 {
	//	return "", errors.New("token invalid")
	//}
	//realToken := tokenArray[1]
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			//responses.Error(w, http.StatusUnauthorized, "ErrSignatureInvalid")
			return "", errors.New("ErrSignatureInvalid")
		}
		//responses.Error(w, http.StatusBadRequest, "Could not parse token")
		return "", errors.New("Could not parse token")
	}
	if !tkn.Valid {
		//responses.Error(w, http.StatusUnauthorized, "Token invalid")
		return "", errors.New("Token Invalid")
	}
	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 60*time.Second {
		//responses.Error(w, 400, "Could not create new token")
		return "", errors.New("Could not create new token")
	}
	expirationTime := time.Now().Add(2 * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()
	NewToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	NewTokenString, err := NewToken.SignedString(jwtKey)
	if err != nil {
		//w.WriteHeader(http.StatusInternalServerError)
		return "", errors.New("StatusInternalServerError")
	}
	return NewTokenString, nil
}
