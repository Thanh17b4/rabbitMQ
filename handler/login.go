package handler

import (
	"encoding/json"
	"fmt"
	model "github.com/Thanh17b4/practice/model"
	"io/ioutil"
	"net/http"
)

type LoginService interface {
	Login(email string, password string) (u *model.User, err error)
	//Activate(code int, email string) (u *model.User, err error)
}
type LoginHandle struct {
	loginService LoginService
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
		fmt.Println("could not marshal your request: ", err.Error())
		return
	}
	user, err := lh.loginService.Login(req.Email, req.Password)
	if err != nil {
		fmt.Println("had an error: ", err.Error())
		json.NewEncoder(w).Encode(map[string]interface{}{
			"massage": err.Error(),
		})
		return
	}
	json.NewEncoder(w).Encode(user)
	return
}
