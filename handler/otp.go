package handler

import (
	"encoding/json"
	"fmt"
	model "github.com/Thanh17b4/practice/model"
	"io/ioutil"
	"net/http"
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
		fmt.Println(" can not marshal your request: ", err.Error())
		return
	}

	userOtp, err := h.otpService.CreatOTPs(userOTP)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"register successfully": userOtp,
	})
}
