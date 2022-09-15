package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Thanh17b4/practice/handler/responses"

	model "github.com/Thanh17b4/practice/model"
)

type UserService interface {
	GetListUser(page int64, limit int64) ([]*model.User, error)
	GetDetailUser(userID int64) (*model.User, error)
	UpdateUserService(user *model.User) (*model.User, error)
	DeleteUser(userID int64) (int64, error)
	CreateUser(user *model.User) (*model.User, error)
}
type UserHandle struct {
	userService UserService
}

func NewUserHandle(userService UserService) *UserHandle {
	return &UserHandle{userService: userService}
}

func (h UserHandle) CreatUserHandle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var user *model.User
	err := json.Unmarshal(reqBody, &user)
	if err != nil {
		responses.Error(w, r, 400, err, "could not marshal your rq")
		return
	}
	userRes, err := h.userService.CreateUser(user)
	if err != nil {
		responses.Error(w, r, http.StatusInternalServerError, err, "could not creat user, userID invalid")
		return
	}
	responses.Success(w, r, http.StatusCreated, userRes)
}
