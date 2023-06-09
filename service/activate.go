package service

import (
	_ "Thanh17b4/practice/model"
	_ "database/sql"
	"errors"
	"time"
)

type ActivateRepo interface {
	Activate(code int, email string) (int, error)
}
type ActivateService struct {
	a ActivateRepo
	u UserRepo
	o OtpRepo
}

func NewActivate(a ActivateRepo, u UserRepo, o OtpRepo) *ActivateService {
	return &ActivateService{a: a, u: u, o: o}
}
func (l *ActivateService) Activate(code int, email string) (int, error) {
	user, err := l.u.GetUserByEmail(email)
	if err != nil {
		return 0, errors.New("email is not valid")
	}
	userOtp, err := l.o.GetUserOTP(user.ID)
	if err != nil {
		return 0, errors.New("userID is not valid")
	}
	if code != userOtp.OTP {
		return 0, errors.New("OTP is not correct")
	}
	t := time.Now()
	if t.After(userOtp.Expired) {
		return 0, errors.New("OTP was expired")
	}
	return l.a.Activate(code, email)
}
