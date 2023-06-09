package handler

import (
	"Thanh17b4/practice/handler/responses"
	"Thanh17b4/practice/model"
	_ "Thanh17b4/practice/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type OtpService interface {
	CreatOTPs(email string) (*model.UserOTP, error)
}
type OtpHandle struct {
	otpService  OtpService
	userService UserService
}

func NewOtpHandle(otpService OtpService) *OtpHandle {
	return &OtpHandle{otpService: otpService}
}

//func (h OtpHandle) CreatUserOTPHandle(w http.ResponseWriter, r *http.Request) {
//	reqBody, _ := ioutil.ReadAll(r.Body)
//	var userOTP *model.UserOTP
//
//	err1 := json.Unmarshal(reqBody, &userOTP)
//	if err1 != nil {
//		fmt.Println(err1)
//		responses.Error(w, r, 400, err1, "Could not marshal your request")
//		return
//	}
//
//	userOtp, err1 := h.otpService.CreatOTPs(userOTP)
//	if err1 != nil {
//		responses.Error(w, r, http.StatusUnauthorized, err1, "Could not creat userOTP")
//		return
//	}
//	responses.Success(w, r, http.StatusAccepted, userOtp)
//}

func (h OtpHandle) CreatUserOTPHandle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	type Req struct {
		Email string
	}
	req := &Req{}
	err1 := json.Unmarshal(reqBody, req)
	if err1 != nil {
		fmt.Println(err1)
		responses.Error(w, r, 400, err1, "Could not marshal your request")
		return
	}

	userOtp, err1 := h.otpService.CreatOTPs(req.Email)
	if err1 != nil {
		responses.Error(w, r, http.StatusUnauthorized, err1, "Could not creat userOTP")
		return
	}
	responses.Success(w, r, http.StatusAccepted, userOtp)
}
