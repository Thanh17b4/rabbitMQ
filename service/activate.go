package service

import (
	"errors"
	"fmt"
	model "github.com/Thanh17b4/practice/model"
	"time"
)

type ActivateService struct {
	otpRepo  OtpRepo
	userRepo UserRepo
}

func NewActivate(otpRepo OtpRepo, userRepo UserRepo) *ActivateService {
	return &ActivateService{otpRepo: otpRepo, userRepo: userRepo}
}
func (l *ActivateService) Activate(code int, email string) (u *model.User, err error) {
	user := l.userRepo.GetUserByEmail(email)
	fmt.Println("user: ", user)
	if user == nil {
		return nil, errors.New("email is not valid")
	}
	userOtp := l.otpRepo.GetOTP(user.ID)
	//fmt.Println("userOtp: ", userOtp)
	if userOtp == nil {
		return nil, errors.New("userOtp is not available")
	}
	if code != userOtp.OTP {
		fmt.Println("OTP is not correct: ")
		return nil, errors.New("OTP is not correct")
	}
	t := time.Now()
	if t.After(userOtp.Expired) {
		return nil, errors.New("OTP was expired")
	}
	return user, nil
}
