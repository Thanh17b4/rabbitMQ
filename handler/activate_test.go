package handler_test

import (
	"bytes"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"

	"github.com/Thanh17b4/practice/handler"
	"github.com/Thanh17b4/practice/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestActivateHandle_Active(t *testing.T) {
	activateService := new(mocks.ActivateService)
	activateHandler := handler.NewActivateHandle(activateService)
	t.Run("When body invalid", func(t *testing.T) {
		body := bytes.NewBufferString(`"abc": "xyz"`)
		w := httptest.NewRecorder()
		r := httptest.NewRequest("POST", "/users/login/active", body)
		r.Header.Set("Content-Type", "application/json")

		activateHandler.Active(w, r)
		got := w.Body.String()
		want := fmt.Sprint(`{"error":{"code":400, "error":"invalid character ':' after top-level value", "massage":"could not marshal your request"}}`)
		assert.Equal(t, 400, w.Code)
		assert.JSONEq(t, want, got)
	})
	t.Run("When email invalid", func(t *testing.T) {
		body := bytes.NewBufferString(`{"email": "abc",
											"otp": "123456"}`)
		r := httptest.NewRequest("POST", "/users/login/active", body)
		r.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		activateService.On("Activate", mock.Anything, mock.Anything).Return("", errors.New("could not find email in database")).Once()

		activateHandler.Active(w, r)
		assert.Equal(t, 400, w.Code)
		want := fmt.Sprint(`{"error":{"code":400, "error":"could not find email in database", "massage":"Email or OTP is not correct"}}`)
		assert.JSONEq(t, want, w.Body.String())
	})
	t.Run("When otp invalid", func(t *testing.T) {
		body := bytes.NewBufferString(`{"email": "thanhpham2@gmail.com",
											"otp": "1234567"}`)
		r := httptest.NewRequest("POST", "/users/login/active", body)
		r.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		activateService.On("Activate", mock.Anything, mock.Anything).Return("", errors.New("OTP is not correct")).Once()

		activateHandler.Active(w, r)
		assert.Equal(t, 400, w.Code)
		want := fmt.Sprint(`{"error":{"code":400, "error":"OTP is not correct", "massage":"Email or OTP is not correct"}}`)
		assert.JSONEq(t, want, w.Body.String())
	})
	t.Run("When otp was expired", func(t *testing.T) {
		body := bytes.NewBufferString(`{"email": "thanhpham2@gmail.com",
											"otp": "123456"}`)
		r := httptest.NewRequest("POST", "/users/login/active", body)
		r.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		activateService.On("Activate", mock.Anything, mock.Anything).Return("", errors.New("OTP was expired")).Once()

		activateHandler.Active(w, r)
		assert.Equal(t, 400, w.Code)
		want := fmt.Sprint(`{"error":{"code":400, "error":"OTP was expired", "massage":"Email or OTP is not correct"}}`)
		assert.JSONEq(t, want, w.Body.String())
	})
	t.Run("When activated successfully", func(t *testing.T) {
		body := bytes.NewBufferString(`{"email": "thanhpham2@gmail.com",
											"otp": "733432"}`)
		r := httptest.NewRequest("POST", "/users/login/active", body)
		r.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		activateService.On("Activate", mock.Anything, mock.Anything).Return("Welcome", nil).Once()
		activateHandler.Active(w, r)

		assert.Equal(t, 202, w.Code)
		want := fmt.Sprint(`{"data":{"Activate successfully": "Welcome"}}`)
		assert.JSONEq(t, want, w.Body.String())
	})

}
