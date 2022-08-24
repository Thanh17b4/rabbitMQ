package handler_test

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/stretchr/testify/assert"

	"github.com/Thanh17b4/practice/handler"
	"github.com/Thanh17b4/practice/tests/mocks"
)

func TestUserHandle_GetDetailUserHandle(t *testing.T) {
	userService := new(mocks.UserService)
	userHandler := handler.NewUserHandle(userService)

	t.Run("type userID incorrect", func(t *testing.T) {

		req := httptest.NewRequest("GET", "/token/users/a", nil)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		userHandler.GetDetailUserHandle(w, req)

		got := w.Body.String()
		want := fmt.Sprint(`{"error":{"code":400, "error":"strconv.Atoi: parsing \"\": invalid syntax", "massage":"userID must be number"}}`)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.JSONEq(t, want, got)
	})

	t.Run("userID invalid", func(t *testing.T) {

		req := httptest.NewRequest("GET", "/token/users/100", nil)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		userService.On("GetDetailUser", mock.Anything).Return(nil, errors.New("could not get user")).Once()
		userHandler.GetDetailUserHandle(w, req)

		got := w.Body.String()
		want := fmt.Sprint(`{"error":{"code":400, "error":"userID is not correct: sql: no rows in result set", "massage":"could not get user"}}`)

		fmt.Println("aa:", w.Body.String())
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.JSONEq(t, want, got)

	})
	t.Run("userID invalid", func(t *testing.T) {

	})
}
