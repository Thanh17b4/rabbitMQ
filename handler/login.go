package handler

import (
	"encoding/json"
	"github.com/Thanh17b4/practice/responses"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"net/http"
	"strings"
)

type LoginService interface {
	Login(email string, password string) (string, error)
	Refresh(token string) (string, error)
}
type LoginHandle struct {
	loginService LoginService
}
type Claim struct {
	Username string
	jwt.StandardClaims
}

func NewLoginHandle(loginService LoginService) *LoginHandle {
	return &LoginHandle{loginService: loginService}
}

func (lh LoginHandle) Login(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	type Req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	req := &Req{}
	err := json.Unmarshal(reqBody, req)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, "could not marshal your request")
		return
	}
	token, err := lh.loginService.Login(req.Email, req.Password)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, "Email or Password are not correct")
		return
	}
	responses.Success(w, map[string]interface{}{
		"token": token,
	})
	return
}

func (lh LoginHandle) Refresh(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	tokenArray := strings.Split(token, " ")
	if len(tokenArray) != 2 {
		responses.Error(w, http.StatusUnauthorized, "currentToken invalid")
		return
	}
	realToken := tokenArray[1]
	newToken, err := lh.loginService.Refresh(realToken)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, "Could not refresh token")
	}
	responses.Success(w, newToken)
}
