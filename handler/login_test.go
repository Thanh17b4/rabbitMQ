package handler_test

import (
	"bytes"
	"context"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/pkg/errors"

	"github.com/stretchr/testify/assert"

	"github.com/Thanh17b4/practice/handler"
	"github.com/Thanh17b4/practice/tests/mocks"
)

func TestLoginHandle_Login(t *testing.T) {
	loginService := new(mocks.LoginService)
	loginHandler := handler.NewLoginHandle(loginService)
	t.Run("When body invalid", func(t *testing.T) {
		body := bytes.NewBufferString(`"abc": "xyz"`)
		w := httptest.NewRecorder()
		r := httptest.NewRequest("POST", "/users/login", body)
		r.Header.Set("Content-Type", "application/json")

		loginHandler.Login(w, r)
		got := w.Body.String()
		want := fmt.Sprint(`{"error":{"code":401, "error":"invalid character ':' after top-level value", "massage":"could not marshal your request"}}`)
		assert.Equal(t, 401, w.Code)
		assert.JSONEq(t, want, got)
	})
	t.Run("When email invalid", func(t *testing.T) {
		body := bytes.NewBufferString(`{"email": "abc",
											"password": "22121992Th"}`)
		req := httptest.NewRequest("POST", "/users/login", body)
		req.Header.Set("Content-Type", "application/json")

		req = req.WithContext(context.Background())
		w := httptest.NewRecorder()

		loginService.On("Login", mock.Anything, mock.Anything).Return("", errors.New("could not find email in database")).Once()

		loginHandler.Login(w, req)
		assert.Equal(t, 401, w.Code)
		want := fmt.Sprint(`{"error":{"code":401, "error":"could not find email in database", "massage":"Login failed: could not find email in database"}}`)
		assert.JSONEq(t, want, w.Body.String())
	})
	t.Run("When password invalid", func(t *testing.T) {
		body := bytes.NewBufferString(`{"email": "thanhpham94@gmail.com",
											"password": "22121992T"}`)
		req := httptest.NewRequest("POST", "/users/login", body)
		req.Header.Set("Content-Type", "application/json")

		req = req.WithContext(context.Background())
		w := httptest.NewRecorder()

		loginService.On("Login", mock.Anything, mock.Anything).Return("", errors.New("password is not correct")).Once()

		loginHandler.Login(w, req)
		assert.Equal(t, 401, w.Code)
		want := fmt.Sprint(`{"error":{"code":401, "error":"password is not correct", "massage":"Login failed: password is not correct"}}`)
		assert.JSONEq(t, want, w.Body.String())
	})
	t.Run("When email and password valid", func(t *testing.T) {
		body := bytes.NewBufferString(`{"email": "thanhpham94@gmail.com",
											"password": "22121992Th"}`)
		req := httptest.NewRequest("POST", "/users/login", body)
		req.Header.Set("Content-Type", "application/json")

		req = req.WithContext(context.Background())
		w := httptest.NewRecorder()

		loginService.On("Login", mock.Anything, mock.Anything).Return("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VybmFtZSI6IiIsIkVtYWlsIjoidGhhbmhwaGFtOTRAZ21haWwuY29tIiwiZXhwIjoxNjYxOTM1MTQ3LCJpYXQiOjE2NjE5MzUwMjd9.l_YcoUUYV_ylu2afmC-oYOLle4DlbyNQL0OTa9zzyZg", nil).Once()

		loginHandler.Login(w, req)
		assert.Equal(t, 200, w.Code)
		got := w.Body.String()
		want := fmt.Sprint(`{"data":{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VybmFtZSI6IiIsIkVtYWlsIjoidGhhbmhwaGFtOTRAZ21haWwuY29tIiwiZXhwIjoxNjYxOTM1MTQ3LCJpYXQiOjE2NjE5MzUwMjd9.l_YcoUUYV_ylu2afmC-oYOLle4DlbyNQL0OTa9zzyZg"}}`)
		assert.JSONEq(t, want, got)
	})

}

//func TestLoginHandle_Refresh(t *testing.T) {
//	loginService := new(mocks.LoginService)
//	loginHandler := handler.NewLoginHandle(loginService)
//	t.Run("When missing token", func(t *testing.T) {
//		w := httptest.NewRecorder()
//		r := httptest.NewRequest("POST", "/token/users/refresh", nil)
//		r.Header.Set("Content-Type", "application/json")
//
//		//loginService.On("Refresh", mock.Anything).Once().Return("", errors.New("Could  not parse token"))
//		loginHandler.Refresh(w, r)
//		assert.Equal(t, 200, w.Code)
//		got := w.Body.String()
//		fmt.Println("aa", got)
//		want := fmt.Sprint(`{"data":{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VybmFtZSI6IiIsIkVtYWlsIjoidGhhbmhwaGFtOTRAZ21haWwuY29tIiwiZXhwIjoxNjYxOTM1MTQ3LCJpYXQiOjE2NjE5MzUwMjd9.l_YcoUUYV_ylu2afmC-oYOLle4DlbyNQL0OTa9zzyZg"}}`)
//		assert.JSONEq(t, want, got)
//
//	})
//}
