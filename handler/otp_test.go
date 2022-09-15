package handler_test

import (
	"bytes"
	"fmt"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Thanh17b4/practice/model"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"

	"github.com/stretchr/testify/assert"

	"github.com/Thanh17b4/practice/handler"
	"github.com/Thanh17b4/practice/tests/mocks"
)

func TestOtpHandle_CreatUserOTPHandle(t *testing.T) {
	otpService := new(mocks.OtpService)
	otpHandler := handler.NewOtpHandle(otpService)
	t.Run("when body input is not correct", func(t *testing.T) {
		body := bytes.NewBufferString(`"userID":"xyz"`)
		w := httptest.NewRecorder()
		r := httptest.NewRequest("POST", "/users_otp/register/otp", body)
		r.Header.Set("Content-Type", "application/json")
		otpHandler.CreatUserOTPHandle(w, r)
		got := w.Body.String()
		want := fmt.Sprint(`{"error":{"code":400, "error":"invalid character ':' after top-level value", "massage":"Could not marshal your request"}}`)
		assert.Equal(t, 400, w.Code)
		assert.JSONEq(t, want, got)
	})
	t.Run("when body input invalid", func(t *testing.T) {
		body := bytes.NewBufferString(`{"user_id": "100"}`)
		req := httptest.NewRequest("POST", "/users_otp/register/otp", body)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		otpService.On("CreatOTPs", mock.Anything).Return(nil, errors.New("userID is not valid: userID is not correct: sql: no rows in result set")).Once()
		otpHandler.CreatUserOTPHandle(w, req)

		got := w.Body.String()
		want := fmt.Sprint(`{"error":{"code":401, "error":"userID is not valid: userID is not correct: sql: no rows in result set", "massage":"Could not creat userOTP"}}`)
		assert.Equal(t, 401, w.Code)
		assert.JSONEq(t, want, got)
	})
	t.Run("when creat userOtp successfully", func(t *testing.T) {
		body := bytes.NewBufferString(`{"user_id": "31"}`)
		req := httptest.NewRequest("POST", "/users_otp/register/otp", body)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		otpService.On("CreatOTPs", mock.Anything).Return(&model.UserOTP{
			UserID:  32,
			OTP:     123456,
			Expired: time.Time{},
			CreatAt: time.Time{},
		}, nil).Once()
		otpHandler.CreatUserOTPHandle(w, req)

		got := w.Body.String()
		want := fmt.Sprint(`{"data":{"creatAt":"0001-01-01T00:00:00Z", "expired":"0001-01-01T00:00:00Z", "otp":123456, "userID":32}}`)
		assert.Equal(t, 202, w.Code)
		assert.JSONEq(t, want, got)
	})
}
