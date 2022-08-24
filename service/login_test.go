package service_test

import (
	"testing"

	"github.com/Thanh17b4/practice/model"

	"github.com/Thanh17b4/practice/service"
	"github.com/Thanh17b4/practice/tests/mocks"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLoginService_Login(t *testing.T) {
	userRepo := new(mocks.UserRepo)
	otpRepo := new(mocks.OtpRepo)
	loginService := service.NewLogin(otpRepo, userRepo)

	t.Run("get user by email get failed", func(t *testing.T) {
		userRepo.On("GetUserByEmail", mock.Anything, mock.Anything).Once().Return(nil, errors.New("could not find email in database"))
		var mail, password string
		got, err := loginService.Login(mail, password)
		assert.Error(t, err)
		assert.Zero(t, got)
		assert.Equal(t, "could not find email in database", err.Error())
	})
	t.Run("hash password get failed", func(t *testing.T) {
		user := &model.User{Email: "thanhpham@gmail.com",
			Password: "$2a$10$8cB5XBRK/ukpnVikGY8e8eQ0LkJ/QxtAQp9rqSYG7uS35GE/lELEi"}
		email := "thanhpham@gmail.com"
		password := "wrong pass"
		_, err := loginService.CompareHashAndPassword(user.Password, password)
		userRepo.On("GetUserByEmail", mock.Anything, mock.Anything).Once().Return(nil, errors.New("password is not correct"))
		got, err := loginService.Login(email, password)
		assert.Error(t, err)
		assert.Equal(t, "", got)
	})
	t.Run("user has not activate", func(t *testing.T) {
		user := &model.User{Email: "thanhpham@gmail.com",
			Password: "$2a$10$8cB5XBRK/ukpnVikGY8e8eQ0LkJ/QxtAQp9rqSYG7uS35GE/lELEi"}
		email := "thanhpham@gmail.com"
		password := "22121992Th"
		_, err := loginService.CompareHashAndPassword(user.Password, password)
		userRepo.On("GetUserByEmail", mock.Anything, mock.Anything).Once().Return(user, nil)
		got, err := loginService.Login(email, password)
		assert.Error(t, err)
		assert.Equal(t, "", got)
	})
	t.Run("create token get fail", func(t *testing.T) {
		user := &model.User{Password: "$2a$10$8cB5XBRK/ukpnVikGY8e8eQ0LkJ/QxtAQp9rqSYG7uS35GE/lELEi",
			Email:     "thanhpv@gmail.com",
			Activated: 1}
		email := "thanhpv@gmail.com"
		password := "222121992Th"
		userRepo.On("GetUserByEmail", mock.Anything, mock.Anything).Once().Return(user, nil)
		_, err := loginService.CreateToken(email)
		got, err := loginService.Login(email, password)
		assert.Error(t, err)
		assert.Equal(t, "", got)
	})
	t.Run("create token get success", func(t *testing.T) {
		user := &model.User{Password: "$2a$10$8cB5XBRK/ukpnVikGY8e8eQ0LkJ/QxtAQp9rqSYG7uS35GE/lELEi",
			Activated: 1}
		var email string
		password := "22121992Th"
		_, err := loginService.CompareHashAndPassword(user.Password, password)
		userRepo.On("GetUserByEmail", mock.Anything, mock.Anything).Once().Return(user, nil)
		loginService.Activate(user.Activated)
		tkn, _ := loginService.CreateToken(email)
		got, err := loginService.Login(email, password)
		assert.Nil(t, err)
		assert.Equal(t, tkn, got)
	})
}
