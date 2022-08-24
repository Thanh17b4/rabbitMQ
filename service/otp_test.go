package service_test

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/Thanh17b4/practice/service"

	"github.com/Thanh17b4/practice/model"
	"github.com/stretchr/testify/mock"

	"github.com/Thanh17b4/practice/tests/mocks"
)

func TestOtpService_CreatOTPs(t *testing.T) {
	userRepo := new(mocks.UserRepo)
	otpRepo := new(mocks.OtpRepo)

	t.Run("err: DetailUser got failed", func(t *testing.T) {
		otpService := service.NewOtpService(userRepo, otpRepo)
		userOTP := model.UserOTP{
			UserID:  4,
			OTP:     100070,
			Expired: time.Now().Add(10000),
			CreatAt: time.Now(),
		}
		userRepo.On("DetailUser", mock.Anything).Once().Return(nil, errors.New("test error"))
		got, err := otpService.CreatOTPs(&userOTP)
		assert.Error(t, err)
		assert.Nil(t, got)
		assert.Equal(t, "userID is not valid: test error", err.Error())
	})

	t.Run("CreatOTP got failed", func(t *testing.T) {
		otpService := service.NewOtpService(userRepo, otpRepo)
		userOTP := model.UserOTP{
			UserID:  4,
			OTP:     100070,
			Expired: time.Now().Add(10000),
			CreatAt: time.Now(),
		}
		userRepo.On("DetailUser", mock.Anything).Once().Return(&model.User{
			Email: "chipv.bka@gmail.com",
			Name:  "Test",
		}, nil)
		otpRepo.On("CreatOTP", mock.Anything).Once().Return(nil, errors.New("test failed"))

		got, err := otpService.CreatOTPs(&userOTP)
		assert.Error(t, err)
		assert.Nil(t, got)
		assert.Equal(t, "test failed", err.Error())
	})
	t.Run("CreatOTP got success", func(t *testing.T) {
		otpService := service.NewOtpService(userRepo, otpRepo)
		userOTP := model.UserOTP{
			UserID:  4,
			OTP:     100070,
			Expired: time.Now().Add(10000),
			CreatAt: time.Now(),
		}
		userRepo.On("DetailUser", mock.Anything).Once().Return(&model.User{
			Email: "chipv.bka@gmail.com",
			Name:  "Test",
		}, nil)
		otpRepo.On("CreatOTP", mock.Anything).Once().Return(&model.UserOTP{}, nil)

		got, err := otpService.CreatOTPs(&userOTP)
		assert.Nil(t, err)
		assert.Equal(t, &model.UserOTP{}, got)
	})

}

var (
	userRepo = new(mocks.UserRepo)
	otpRepo  = new(mocks.OtpRepo)
)
var otpService = service.NewOtpService(userRepo, otpRepo)

func TestOtpService_GetOTPs(t *testing.T) {
	t.Run("Get OTP get failed", func(t *testing.T) {
		otpService := service.NewOtpService(userRepo, otpRepo)
		otpRepo.On("GetUserOTP", mock.Anything).Once().Return(nil, errors.New("had an test error"))
		userOtp := model.UserOTP{
			UserID: 5,
		}
		got, err := otpService.GetOTPs(userOtp.UserID)
		assert.Nil(t, got)
		assert.Error(t, err)
		assert.Equal(t, "had an test error", err.Error())
	})

	t.Run("Get OTP get pass", func(t *testing.T) {
		otpRepo.On("GetUserOTP", mock.Anything).Once().Return(&model.UserOTP{}, nil)
		userOtp := &model.UserOTP{UserID: 5}
		got, err := otpService.GetOTPs(userOtp.UserID)
		assert.Nil(t, err)
		assert.Equal(t, &model.UserOTP{}, got)
	})
}
