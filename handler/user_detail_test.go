package handler_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Thanh17b4/practice/handler"

	"github.com/Thanh17b4/practice/model"

	"github.com/go-chi/chi/v5"

	"github.com/stretchr/testify/mock"

	"github.com/stretchr/testify/assert"

	"github.com/Thanh17b4/practice/tests/mocks"
)

func TestUserHandle_GetDetailUserHandle(t *testing.T) {
	userService := new(mocks.UserService)
	userHandler := handler.NewUserHandle(userService)

	t.Run("type userID incorrect", func(t *testing.T) {

		req := httptest.NewRequest("GET", "/token/users/a", nil)
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("id", "a")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))

		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		userHandler.GetDetailUserHandle(w, req)

		got := w.Body.String()
		want := fmt.Sprint(`{"error":{"code":400, "error":"strconv.ParseInt: parsing \"a\": invalid syntax", "massage":"userID must be number"}}`)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.JSONEq(t, want, got)
	})

	t.Run("userID invalid", func(t *testing.T) {

		req := httptest.NewRequest("GET", "/token/users/100", nil)

		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("id", "100")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))

		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		userService.On("GetDetailUser", mock.Anything).Return(nil, errors.New("userID is not correct: sql: no rows in result set")).Once()
		userHandler.GetDetailUserHandle(w, req)

		got := w.Body.String()
		want := fmt.Sprint(`{"error":{"code":400, "error":"userID is not correct: sql: no rows in result set", "massage":"could not get user"}}`)

		fmt.Println("aa:", w.Body.String())
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.JSONEq(t, want, got)

	})
	t.Run("userID success", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/token/users/36", nil)
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("id", "36")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		userService.On("GetDetailUser", mock.Anything).Return(&model.User{
			ID:        36,
			Name:      "Thanhabc1",
			Email:     "thanhpham96@gmail.com",
			Protected: 0,
			Banned:    0,
			Activated: 0,
			Address:   "China",
			Password:  "",
			Username:  "thanh24",
		}, nil)
		userHandler.GetDetailUserHandle(w, req)

		got := w.Body.String()
		want := fmt.Sprint(`{"data":{"id":36,"name":"Thanhabc1","email":"thanhpham96@gmail.com","protected":0,"banned":0,"activated":0,"address":"China","password":"","username":"thanh24"}}`)
		fmt.Println("bb:", w.Body.String())
		assert.Equal(t, http.StatusAccepted, w.Code)
		assert.JSONEq(t, want, got)
	})
}
