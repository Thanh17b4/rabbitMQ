package service_test

import (
	"errors"
	"testing"
	"time"

	"github.com/Thanh17b4/practice/model"

	"github.com/Thanh17b4/practice/service"
	"github.com/Thanh17b4/practice/tests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestActivateService_Activate(t *testing.T) {
	otpRepo := new(mocks.OtpRepo)
	userRepo := new(mocks.UserRepo)
	activateService := service.NewActivate(otpRepo, userRepo)
	t.Run("get user by email failed", func(t *testing.T) {
		userRepo.On("GetUserByEmail", mock.Anything).Once().Return(nil, errors.New("email is not valid"))
		var (
			email = "abc@gmail.com"
			code  int
		)
		got, err := activateService.Activate(code, email)
		assert.Error(t, err)
		assert.Equal(t, "", got)
	})
	t.Run("get userOTP failed", func(t *testing.T) {
		user := &model.User{
			Email: "thanhpham@gmail.com",
			ID:    0,
		}
		userRepo.On("GetUserByEmail", mock.Anything).Once().Return(user, nil)
		otpRepo.On("GetUserOTP", mock.Anything).Once().Return(nil, errors.New("userID is not valid"))
		var code int
		got, err := activateService.Activate(code, user.Email)
		assert.Error(t, err)
		assert.Equal(t, "", got)
	})
	t.Run("Enter OTP failed", func(t *testing.T) {
		user := &model.User{
			Email: "thanhpham@gmail.com",
			ID:    47,
		}
		userRepo.On("GetUserByEmail", mock.Anything).Once().Return(user, nil)
		otpRepo.On("GetUserOTP", mock.Anything).Once().Return(&model.UserOTP{}, nil)
		OTP := 3000
		got, err := activateService.Activate(OTP, user.Email)
		assert.Error(t, err)
		assert.Equal(t, "", got)
	})
	t.Run("OTP expired", func(t *testing.T) {
		user := &model.User{
			Email: "thanhpham@gmail.com",
			ID:    47,
		}
		userRepo.On("GetUserByEmail", mock.Anything).Once().Return(user, nil)
		userOtp := &model.UserOTP{UserID: user.ID,
			OTP:     123456,
			Expired: time.Now().Add(time.Second * (-10))}
		otpRepo.On("GetUserOTP", mock.Anything).Once().Return(userOtp, nil)
		OTP := 123456
		got, err := activateService.Activate(OTP, user.Email)
		assert.Error(t, err)
		assert.Equal(t, "", got)
	})
	t.Run("OTP expired", func(t *testing.T) {
		user := &model.User{
			Email: "thanhpham@gmail.com",
			ID:    47,
		}
		userRepo.On("GetUserByEmail", mock.Anything).Once().Return(user, nil)
		userOtp := &model.UserOTP{UserID: user.ID,
			OTP:     123456,
			Expired: time.Now().Add(time.Second * 10)}
		otpRepo.On("GetUserOTP", mock.Anything).Once().Return(userOtp, nil)
		OTP := 123456
		got, err := activateService.Activate(OTP, user.Email)
		assert.Nil(t, err)
		assert.Equal(t, "Welcome", got)
	})
}
