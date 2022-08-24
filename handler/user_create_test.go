package handler_test

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Thanh17b4/practice/model"

	"github.com/stretchr/testify/mock"

	"github.com/stretchr/testify/assert"

	"github.com/Thanh17b4/practice/handler"
	"github.com/Thanh17b4/practice/tests/mocks"
)

func TestUserHandle_CreatUserHandle(t *testing.T) {
	userService := new(mocks.UserService)
	userHandler := handler.NewUserHandle(userService)
	t.Run("marshal request get failed", func(t *testing.T) {
		body := bytes.NewBufferString(`"test": "abc-xyz"`)
		req := httptest.NewRequest("POST", "/todos", body)
		req.Header.Set("Content-Type", "application/json")
		ctx := req.Context()
		req = req.WithContext(ctx)
		w := httptest.NewRecorder()

		userHandler.CreatUserHandle(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		//userService.AssertExpectations(t)
		expectedResponse := fmt.Sprint(`{"error":{"code":400, "error":"invalid character ':' after top-level value", "massage":"could not marshal your rq"}}`)
		fmt.Println("aa:", w.Body.String())
		assert.JSONEq(t, expectedResponse, w.Body.String())
	})
	t.Run("Create user failed", func(t *testing.T) {
		body := bytes.NewBufferString(`{"title": "task1"}`)
		req := httptest.NewRequest("POST", "/todos", body)
		req.Header.Set("Content-Type", "application/json")

		req = req.WithContext(context.Background())
		w := httptest.NewRecorder()

		userService.On("CreateUser", mock.Anything).Return(nil, errors.New("could not create user")).Once()

		userHandler.CreatUserHandle(w, req)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		userService.AssertExpectations(t)
		expectedResponse := fmt.Sprint(`{"error":{"code":500, "error":"could not create user", "massage":"could not creat user"}}`)
		assert.JSONEq(t, expectedResponse, w.Body.String())
	})
	t.Run("Create user success", func(t *testing.T) {
		body := bytes.NewBufferString(`{
			"email": "thanhpham2@gmail.com",
			"name": "Thanhabc1",
			"address": "China",
			"password": "22121992T",
			"activated": 0
		}`)
		req := httptest.NewRequest("POST", "/users/register", body)
		req.Header.Set("Content-Type", "application/json")
		req = req.WithContext(context.Background())
		w := httptest.NewRecorder()

		userService.On("CreateUser", mock.Anything).Return(&model.User{
			ID:        100,
			Name:      "Thanhabc1",
			Email:     "thanhpham2@gmail.com",
			Protected: 0,
			Banned:    0,
			Activated: 0,
			Address:   "nil",
		}, nil).Once()

		userHandler.CreatUserHandle(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)
		userService.AssertExpectations(t)
		expectedResponse := fmt.Sprintln(`{"data":{"id":100,"name":"Thanhabc1","email":"thanhpham2@gmail.com","protected":0,"banned":0,"activated":0,"address":"nil","password":"","username":""}}`)
		fmt.Println("bb:", w.Body.String())
		assert.JSONEq(t, expectedResponse, w.Body.String())
	})
}
