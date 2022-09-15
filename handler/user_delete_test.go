package handler

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	//"github.com/Thanh17b4/practice/model"
	"github.com/Thanh17b4/practice/tests/mocks"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserHandle_DeleteUserHandle(t *testing.T) {
	userService := new(mocks.UserService)
	userHandler := NewUserHandle(userService)

	t.Run("type userID incorrect", func(t *testing.T) {

		req := httptest.NewRequest("GET", "/token/users/a", nil)
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("id", "a")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))

		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		userHandler.DeleteUserHandle(w, req)

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

		userService.On("DeleteUser", mock.Anything).Return(int64(0), errors.New("userID is not correct: sql: no rows in result set")).Once()
		userHandler.DeleteUserHandle(w, req)

		got := w.Body.String()
		want := fmt.Sprint(`{"error":{"code":400, "error":"userID is not correct: sql: no rows in result set", "massage":"Could not delete user"}}`)

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

		userService.On("DeleteUser", mock.Anything).Return(int64(36), nil)
		userHandler.DeleteUserHandle(w, req)

		got := w.Body.String()
		want := fmt.Sprint(`{"data":"userID 36 has been deleted"}`)
		fmt.Println("bb:", w.Body.String())
		assert.Equal(t, http.StatusAccepted, w.Code)
		assert.JSONEq(t, want, got)
	})
}
