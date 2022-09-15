package handler_test

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Thanh17b4/practice/model"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"

	"github.com/go-chi/chi/v5"

	"github.com/Thanh17b4/practice/handler"

	"github.com/Thanh17b4/practice/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestUserHandle_UpdateUserHandle(t *testing.T) {
	userService := new(mocks.UserService)
	userHandler := handler.NewUserHandle(userService)

	t.Run("type of id is not correct", func(t *testing.T) {
		body := bytes.NewBufferString(`"email": "thanhpham.hvnh@gmail.com",
		"password": "22121992",
		"name": "Thanh",
		"address": "China",
		"username": "thanh26"`)
		r := httptest.NewRequest(http.MethodPut, "/token/users/a", body)
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("id", "a")
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, chiCtx))
		w := httptest.NewRecorder()

		userHandler.UpdateUserHandle(w, r)

		got := w.Body.String()
		want := fmt.Sprint(`{"error":{"code":400, "error":"strconv.ParseInt: parsing \"a\": invalid syntax", "massage":"userID must be number"}}`)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.JSONEq(t, want, got)
	})
	t.Run("marshal request get failed", func(t *testing.T) {
		body := bytes.NewBufferString(`"email": "thanhpham.hvnh@gmail.com",
		"password": "22121992",
		"name": "Thanh",
		"address": "China",
		"username": "thanh26"`)
		r := httptest.NewRequest("PUT", "/token/users/100", body)
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("id", "100")
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, chiCtx))
		w := httptest.NewRecorder()
		//userService.On("UpdateUserService", mock.Anything).Once().Return(nil, errors.New("abc"))
		userHandler.UpdateUserHandle(w, r)

		got := w.Body.String()
		want := fmt.Sprint(`{"error":{"code":400, "error":"invalid character ':' after top-level value", "massage":"could not Unmarshal body request"}}`)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.JSONEq(t, want, got)
	})
	t.Run("When id invalid", func(t *testing.T) {
		body := bytes.NewBufferString(`{"email": "thanhpham.hvnh@gmail.com",
		"password": "22121992",
		"name": "Thanh",
		"address": "China",
		"username": "thanh26"}`)
		r := httptest.NewRequest("PUT", "/token/users/100", body)
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("id", "100")
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, chiCtx))
		w := httptest.NewRecorder()
		userService.On("UpdateUserService", mock.Anything).Once().Return(nil, errors.New("userID is not correct: sql: no rows in result set"))
		userHandler.UpdateUserHandle(w, r)

		got := w.Body.String()
		want := fmt.Sprint(`{"error":{"code":400, "error":"userID is not correct: sql: no rows in result set", "massage":"Could not update user because information input is not correct"}}`)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.JSONEq(t, want, got)
	})

	t.Run("When missing require param", func(t *testing.T) {
		body := bytes.NewBufferString(`{"email": "thanhpham.hvnh@gmail.com",
		"password": "22121992",
		"name": "Thanh",
		"address": "China",
		"username": "thanh26"}`)
		r := httptest.NewRequest("PUT", "/token/users/1", body)
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("id", "1")
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, chiCtx))
		w := httptest.NewRecorder()
		userService.On("UpdateUserService", mock.Anything).Once().Return(nil, errors.New("required field can not empty"))
		userHandler.UpdateUserHandle(w, r)

		got := w.Body.String()
		want := fmt.Sprint(`{"error":{"code":400, "error":"required field can not empty", "massage":"Could not update user because information input is not correct"}}`)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.JSONEq(t, want, got)
	})

	t.Run("When updated successfully", func(t *testing.T) {
		body := bytes.NewBufferString(`{"email": "thanhpham.hvnh@gmail.com",
		"password": "22121992",
		"name": "Thanh",
		"address": "China",
		"username": "thanh26"}`)
		r := httptest.NewRequest("PUT", "/token/users/1", body)
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("id", "1")
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, chiCtx))
		w := httptest.NewRecorder()
		userService.On("UpdateUserService", mock.Anything).Once().Return(&model.User{
			ID:        1,
			Name:      "Thanh",
			Email:     "thanhpham.hvnh@gmail.com",
			Protected: 0,
			Banned:    0,
			Activated: 0,
			Address:   "China",
			Password:  "22121992",
			Username:  "thanh26",
		}, nil)
		userHandler.UpdateUserHandle(w, r)

		got := w.Body.String()
		want := fmt.Sprint(`{"data":{"activated":0, "address":"China", "banned":0, "email":"thanhpham.hvnh@gmail.com", "id":1, "name":"Thanh", "password":"22121992", "protected":0, "username":"thanh26"}}`)
		assert.Equal(t, 202, w.Code)
		assert.JSONEq(t, want, got)
	})

}
