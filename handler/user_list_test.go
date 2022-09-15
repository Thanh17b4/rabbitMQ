package handler_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Thanh17b4/practice/handler"

	"github.com/Thanh17b4/practice/model"

	"github.com/pkg/errors"

	"github.com/stretchr/testify/mock"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"

	"github.com/Thanh17b4/practice/tests/mocks"
)

func TestUserHandle_GetListUserHandle(t *testing.T) {
	userService := new(mocks.UserService)
	userHandler := handler.NewUserHandle(userService)
	t.Run("page type is not correct", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/token/users", nil)

		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("page", "a")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
		w := httptest.NewRecorder()
		userHandler.GetListUser(w, req)
		got := w.Body.String()
		want := fmt.Sprint(`{"error":{"code":400,"massage":"page must be number","error":"strconv.ParseInt: parsing \"\": invalid syntax"}}`)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.JSONEq(t, want, got)
	})
	t.Run("page correct, perPage failed", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/token/users?page=1&perPage=a", nil)
		chiCtx := chi.NewRouteContext()
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
		type Params map[string]string
		params := Params{
			"page":    "1",
			"perPage": "a",
		}
		for key, value := range params {
			chiCtx.URLParams.Add(key, value)
		}
		w := httptest.NewRecorder()

		userHandler.GetListUser(w, req)
		got := w.Body.String()
		want := fmt.Sprint(`{"error":{"code":400,"massage":"perPage must be number","error":"strconv.ParseInt: parsing \"a\": invalid syntax"}}
`)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.JSONEq(t, want, got)
	})
	t.Run("page or limit is zero", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/token/users?page=0&perPage=0", nil)
		chiCtx := chi.NewRouteContext()
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, chiCtx))
		type Params map[string]string
		params := Params{
			"page":    "0",
			"perPage": "0",
		}
		for key, value := range params {
			chiCtx.URLParams.Add(key, value)
		}
		userService.On("GetListUser", mock.Anything, mock.Anything).Once().Return(nil, errors.New("page and limit must be bigger than 0"))
		userHandler.GetListUser(w, r)
		got := w.Body.String()
		fmt.Println("aa: ", got)
		want := fmt.Sprint(`{"error":{"code":400, "error":"page and limit must be bigger than 0", "massage":"Could not get list users"}}`)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.JSONEq(t, want, got)
	})
	t.Run("page is not exist", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/token/users?page=100&perPage=0", nil)
		chiCtx := chi.NewRouteContext()
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, chiCtx))
		type Params map[string]string
		params := Params{
			"page":    "100",
			"perPage": "10",
		}
		for key, value := range params {
			chiCtx.URLParams.Add(key, value)
		}
		userService.On("GetListUser", mock.Anything, mock.Anything).Once().Return(nil, errors.New("Number of pages is too large, page is not exist"))
		userHandler.GetListUser(w, r)
		got := w.Body.String()
		want := fmt.Sprint(`{"error":{"code":400, "error":"Number of pages is too large, page is not exist", "massage":"Could not get list users"}}`)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.JSONEq(t, want, got)
	})
	t.Run("get list users successfully", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/token/users?page=1&perPage=10", nil)
		chiCtx := chi.NewRouteContext()
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, chiCtx))
		type Params map[string]string
		params := Params{
			"page":    "1",
			"perPage": "10",
		}
		for key, value := range params {
			chiCtx.URLParams.Add(key, value)
		}
		user1 := &model.User{
			ID:        100,
			Name:      "Test",
			Email:     "tests@gmail.com",
			Protected: 0,
			Banned:    0,
			Activated: 0,
			Address:   "18 Ton that Thuyet",
			Password:  "",
			Username:  "tests",
		}

		var listUsers []*model.User
		listUsers = append(listUsers, user1)

		userService.On("GetListUser", mock.Anything, mock.Anything).Once().Return(listUsers, nil)
		userHandler.GetListUser(w, r)
		got := w.Body.String()
		want := fmt.Sprint(`{"data":{"id":100,"name":"Test","email":"tests@gmail.com","protected":0,"banned":0,"activated":0,"address":"18 Ton that Thuyet","password":"","username":"tests"}}`)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.JSONEq(t, want, got)
	})
}
