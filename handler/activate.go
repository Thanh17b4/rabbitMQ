package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Thanh17b4/practice/handler/responses"
)

type ActivateService interface {
	Activate(code int, email string) (u string, err error)
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
		responses.Error(w, r, http.StatusBadRequest, err, "could not marshal your request")
		return
	}
	_, err = lh.activateService.Activate(req.Code, req.Email)
	if err != nil {
		responses.Error(w, r, 400, err, "Email or OTP is not correct")
		return
	}
	responses.Success(w, r, http.StatusAccepted, map[string]interface{}{
		"Activate successfully": "Welcome",
	})
	return
}
