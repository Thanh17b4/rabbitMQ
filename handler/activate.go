package handler

import (
	"Thanh17b4/practice/handler/responses"
	"encoding/json"
	_ "github.com/go-chi/chi/v5"
	"io/ioutil"
	"net/http"
)

type ActivateService interface {
	Activate(code int, email string) (u int, err error)
}
type ActivateHandler struct {
	activateService ActivateService
}

func NewActivateHandler(activateService ActivateService) *ActivateHandler {
	return &ActivateHandler{activateService: activateService}
}
func (lh ActivateHandler) Active(w http.ResponseWriter, r *http.Request) {
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
