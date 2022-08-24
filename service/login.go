package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/Thanh17b4/practice/model"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
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

func (l LoginService) CompareHashAndPassword(hashedPassword, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false, errors.New("password is not correct")
	}
	return true, nil
}

func (l LoginService) CreateToken(email string) (string, error) {
	user := &model.User{}
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

func (l LoginService) Activate(number int) (string, error) {
	if number == 0 {
		return "", errors.New("account has not been activated, please activate your account")
	}
	return "your account has been activated", nil
}

func (l LoginService) Login(email string, password string) (string, error) {
	user, err1 := l.userRepo.GetUserByEmail(email)
	if err1 != nil {
		fmt.Println("kkkk", err1.Error())
		return "", errors.New("could not find email in database")
	}
	_, err2 := l.CompareHashAndPassword(user.Password, password)

	if err2 != nil {
		return "", err2
	}
	msg, err3 := l.Activate(user.Activated)
	if err3 != nil {
		return msg, err3
	}

	token, err4 := l.CreateToken(email)
	if err4 != nil {
		return "", err4
	}

	return token, nil
}

func (l LoginService) Refresh(token string) (string, error) {
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return "", errors.New("ErrSignatureInvalid")
		}
		return "", errors.New("Could  not parse token")
	}
	if !tkn.Valid {
		return "", errors.New("Token  Invalid")
	}
	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 60*time.Second {
		return "", errors.New("You can only refresh token after 60s from login successfully ")
	}
	expirationTime := time.Now().Add(2 * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()
	NewToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	NewTokenString, err := NewToken.SignedString(jwtKey)
	if err != nil {
		return "", errors.New("StatusInternalServerError")
	}
	return NewTokenString, nil
}
