package service_test

import (
	"errors"
	"testing"

	"github.com/Thanh17b4/practice/service"
	"github.com/Thanh17b4/practice/tests/mocks"

	"github.com/Thanh17b4/practice/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserService_CreatUser(t *testing.T) {
	userRepo := new(mocks.UserRepo)
	t.Run("Create user failed", func(t *testing.T) {
		user := model.User{}
		userRepo.On("Create", mock.Anything).Once().Return(nil, errors.New("could not Generate Password"))
		userService := service.NewUserService(userRepo)
		got, err := userService.CreateUser(&user)
		assert.Error(t, err)
		assert.Nil(t, got)
		assert.Equal(t, "could not Generate Password", err.Error())
	})
	t.Run("Create user success", func(t *testing.T) {
		user := model.User{Password: "22121992Th"}
		userRepo.On("Create", mock.Anything).Once().Return(&model.User{}, nil)
		userService := service.NewUserService(userRepo)
		got, err := userService.CreateUser(&user)
		assert.Nil(t, err)
		assert.Equal(t, &model.User{}, got)
	})
}

func TestUserService_GetDetailUser(t *testing.T) {
	userRepo := new(mocks.UserRepo)
	t.Run("Get user fail", func(t *testing.T) {
		userRepo.On("DetailUser", mock.Anything).Once().Return(nil, errors.New("could not get user information"))
		userService := service.NewUserService(userRepo)
		user := &model.User{}
		got, err := userService.GetDetailUser(int64(user.ID))
		assert.Error(t, err)
		assert.Nil(t, got)
		assert.Equal(t, "could not get user information", err.Error())
	})
	t.Run("Get user pass", func(t *testing.T) {
		userRepo.On("DetailUser", mock.Anything).Once().Return(&model.User{}, nil)
		userService := service.NewUserService(userRepo)
		user := &model.User{ID: 5}
		got, err := userService.GetDetailUser(int64(user.ID))
		assert.Nil(t, err)
		assert.Equal(t, &model.User{}, got)
	})
}

func TestUserService_GetListUsers(t *testing.T) {
	userRepo := new(mocks.UserRepo)
	userService := service.NewUserService(userRepo)
	t.Run("Get list user fail", func(t *testing.T) {
		userRepo.On("ListUser", mock.Anything, mock.Anything).Once().Return(nil, errors.New("could not get list user"))
		page, limit := 1, 10
		got, err := userService.GetListUser(int64(page), int64(limit))
		assert.Nil(t, got)
		assert.Error(t, err)
		assert.Equal(t, "could not get list user", err.Error())
	})
	t.Run("Get list user success", func(t *testing.T) {
		userRepo.On("ListUser", mock.Anything, mock.Anything).Once().Return([]*model.User{}, nil)
		page, limit := 1, 10
		got, err := userService.GetListUser(int64(page), int64(limit))
		assert.Nil(t, err)
		assert.Equal(t, []*model.User{}, got)
	})
}

func TestUserService_UpdateUserServiceUsers(t *testing.T) {
	userRepo := new(mocks.UserRepo)
	userService := service.NewUserService(userRepo)
	t.Run("Update user fail", func(t *testing.T) {
		userRepo.On("UpdateUser", mock.Anything).Once().Return(nil, errors.New("could not update user"))
		user := &model.User{}
		got, err := userService.UpdateUserService(user)
		assert.Nil(t, got)
		assert.Error(t, err)
		assert.Equal(t, "could not update user", err.Error())
	})
	t.Run("Get list user success", func(t *testing.T) {
		userRepo.On("UpdateUser", mock.Anything).Once().Return(&model.User{}, nil)
		user := &model.User{}
		got, err := userService.UpdateUserService(user)
		assert.Nil(t, err)
		assert.Equal(t, &model.User{}, got)
	})
}
func TestUserService_DeleteUser(t *testing.T) {
	userRepo := new(mocks.UserRepo)
	userService := service.NewUserService(userRepo)
	t.Run("Delete user get fail", func(t *testing.T) {
		userRepo.On("Delete", mock.Anything).Once().Return(int64(0), errors.New("could not delete user"))
		user := &model.User{ID: 10}
		got, err := userService.DeleteUser(int64(user.ID))
		//assert.Exactly(t, got, int64(0))
		assert.Zero(t, got)
		assert.Error(t, err)
		assert.Equal(t, "could not delete user", err.Error())
	})
	t.Run("Delete user get success", func(t *testing.T) {
		user := &model.User{}
		userRepo.On("Delete", mock.Anything).Once().Return(int64(user.ID), nil)
		got, err := userService.DeleteUser(int64(user.ID))
		assert.Nil(t, err)
		assert.Equal(t, int64(user.ID), got)
	})
}
