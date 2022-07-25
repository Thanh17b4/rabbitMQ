package service

import (
	"errors"
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
	user, err := l.userRepo.GetUserByEmail(email)
	if user == nil {
		return nil, errors.New("email is not valid")
	}
	userOtp, err := l.otpRepo.GetOTP(user.ID)
	if err != nil {
		return nil, errors.New("userID is not valid")
	}
	if code != userOtp.OTP {
		return nil, errors.New("OTP is not correct")
	}
	t := time.Now()
	if t.After(userOtp.Expired) {
		return nil, errors.New("OTP was expired")
	}
	return user, nil
}
