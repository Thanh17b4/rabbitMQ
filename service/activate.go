package service

import (
	"errors"
	"time"
)

type ActivateService struct {
	otpRepo  OtpRepo
	userRepo UserRepo
}

func NewActivate(otpRepo OtpRepo, userRepo UserRepo) *ActivateService {
	return &ActivateService{otpRepo: otpRepo, userRepo: userRepo}
}
func (l *ActivateService) Activate(code int, email string) (string, error) {
	user, err := l.userRepo.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("email is not valid")
	}
	userOtp, err := l.otpRepo.GetUserOTP(user.ID)
	if err != nil {
		return "", errors.New("userID is not valid")
	}
	if code != userOtp.OTP {
		return "", errors.New("OTP is not correct")
	}
	t := time.Now()
	if t.After(userOtp.Expired) {
		return "", errors.New("OTP was expired")
	}
	return "Welcome", nil
}
