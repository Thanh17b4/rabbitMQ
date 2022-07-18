package handler

import (
	"encoding/json"
	"fmt"
	model "github.com/Thanh17b4/practice/model"
	"io/ioutil"
	"net/http"
)

type ActivateService interface {
	Activate(code int, email string) (u *model.User, err error)
}
type ActivateHandle struct {
	activateService ActivateService
}

func NewActivateHandle(activateService ActivateService) *ActivateHandle {
	return &ActivateHandle{activateService: activateService}
}
func (lh ActivateHandle) Active(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	type Req struct {
		Email string `json:"email"`
		Code  int    `json:"code"`
	}
	req := &Req{}
	err := json.Unmarshal(reqBody, req)
	if err != nil {
		fmt.Println("could not marshal your request: ", err.Error())
		return
	}
	user, err := lh.activateService.Activate(req.Code, req.Email)
	fmt.Println("req: ", req.Code, req.Email)
	if err != nil {
		fmt.Println("had an error: ", err.Error())
		json.NewEncoder(w).Encode(map[string]interface{}{
			"massage": err.Error(),
		})
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"massage": "login successfully",
	})
	json.NewEncoder(w).Encode(user)
	return
}
