package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Thanh17b4/practice/handler/responses"
	model "github.com/Thanh17b4/practice/model"
)

type OtpService interface {
	CreatOTPs(otp *model.UserOTP) (*model.UserOTP, error)
}
type OtpHandle struct {
	otpService OtpService
}

func NewOtpHandle(otpService OtpService) *OtpHandle {
	return &OtpHandle{otpService: otpService}
}
func (h OtpHandle) CreatUserOTPHandle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var userOTP *model.UserOTP
	err := json.Unmarshal(reqBody, &userOTP)
	if err != nil {
		responses.Error(w, r, 400, err, "Could not marshal your request")
		return
	}

	userOtp, err := h.otpService.CreatOTPs(userOTP)
	if err != nil {
		responses.Error(w, r, http.StatusUnauthorized, err, "Could not creat userOTP")
		return
	}
	responses.Success(w, r, http.StatusAccepted, userOtp)
}
