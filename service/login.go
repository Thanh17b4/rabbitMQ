package service

import (
	"errors"
	"fmt"
	model "github.com/Thanh17b4/practice/model"
	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	otpRepo  OtpRepo
	userRepo UserRepo
}

func NewLogin(otpRepo OtpRepo, userRepo UserRepo) *LoginService {
	return &LoginService{otpRepo: otpRepo, userRepo: userRepo}
}

func (l LoginService) Login(email string, password string) (u *model.User, err error) {
	user := l.userRepo.GetUserByEmail(email)
	fmt.Println("user: ", user)
	if email != user.Email {
		fmt.Println(" could not find email in database: ", err.Error())
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		fmt.Println("password is not correct: ", err.Error())
		return nil, err
	}

	if user.Activated == 0 {
		err = errors.New("account has not been activated, please activate your account")
		return nil, err
	}

	return user, nil
}
