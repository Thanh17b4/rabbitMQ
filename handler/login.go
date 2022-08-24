package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/Thanh17b4/practice/handler/responses"
	"github.com/dgrijalva/jwt-go"
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
	err1 := json.Unmarshal(reqBody, req)
	if err1 != nil {
		responses.Error(w, r, http.StatusUnauthorized, err1, "could not marshal your request")
		return
	}
	token, err2 := lh.loginService.Login(req.Email, req.Password)
	if err2 != nil {
		responses.Error(w, r, http.StatusUnauthorized, err2, "Login failed: "+err2.Error())
		return
	}
	responses.Success(w, r, http.StatusOK, map[string]interface{}{
		"token": token,
	})
	return
}

func (lh LoginHandle) Refresh(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	tokenArray := strings.Split(token, " ")
	if len(tokenArray) != 2 {
		responses.Error(w, r, http.StatusUnauthorized, nil, "currentToken invalid")
		return
	}
	realToken := tokenArray[1]
	newToken, err := lh.loginService.Refresh(realToken)
	if err != nil {
		responses.Error(w, r, http.StatusUnauthorized, err, "Could not refresh token")
	}
	responses.Success(w, r, http.StatusAccepted, newToken)
}
