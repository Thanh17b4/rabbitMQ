package handler

import (
	"encoding/json"
	model "github.com/Thanh17b4/practice/model"
	"github.com/Thanh17b4/practice/responses"
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
		responses.Error(w, http.StatusBadRequest, "could not marshal your request")
		return
	}
	user, err := lh.activateService.Activate(req.Code, req.Email)
	if err != nil {
		responses.Error(w, 400, "Email or OTP is not correct")
		return
	}
	responses.Success(w, map[string]interface{}{
		"Activate successfully, hello": user.Name,
	})
	return
}
