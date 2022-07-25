package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	model "github.com/Thanh17b4/practice/model"
	"github.com/Thanh17b4/practice/responses"
	"github.com/go-chi/chi/v5"
)

type UserService interface {
	GetListUser(page int64, limit int64) ([]*model.User, error)
	GetDetailUser(userID int64) (*model.User, error)
	UpdateUserService(user *model.User) (*model.User, error)
	DeleteUser(userID int64) (int64, error)
	CreatUser(user *model.User) (*model.User, error)
}
type UserHandle struct {
	userService UserService
}

func NewUserHandle(userService UserService) *UserHandle {
	return &UserHandle{userService: userService}
}
func (h UserHandle) GetListUser(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	currentPage, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		responses.Error(w, 400, "page must be number")
		return
	}
	perPage := r.URL.Query().Get("perPage")
	limit, err := strconv.ParseInt(perPage, 10, 64)
	if err != nil {
		responses.Error(w, 400, "perPage must be number")
		return
	}
	users, err := h.userService.GetListUser(currentPage, limit)
	for _, user := range users {
		responses.Success(w, user)
	}
}

func (h UserHandle) GetDetailUserHandle(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		responses.Error(w, 400, "userID must be number")
		return
	}
	user, err := h.userService.GetDetailUser(id)
	if err != nil {
		responses.Error(w, 400, "could not get user")
		return
	}
	responses.Success(w, user)
}

func (h UserHandle) UpdateUserHandle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var user *model.User
	err := json.Unmarshal(reqBody, &user)
	if err != nil {
		responses.Error(w, 400, "could not Unmarshal body request")
		return
	}
	userID := chi.URLParam(r, "id")

	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		responses.Error(w, 400, "userID must be number")
		return
	}
	user.ID = int(id)
	updateID, err := h.userService.UpdateUserService(user)
	if err != nil {
		responses.Error(w, 400, "")
		return
	}
	responses.Success(w, map[string]interface{}{
		"updatedID": updateID,
	})
}
func (h UserHandle) DeleteUserHandle(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		responses.Error(w, 400, "userID must be number")
		return
	}
	deleteId, _ := h.userService.DeleteUser(id)
	responses.Success(w, deleteId)
}

func (h UserHandle) CreatUserHandle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var user *model.User
	err := json.Unmarshal(reqBody, &user)
	if err != nil {
		responses.Error(w, 400, "could not marshal your rq")
		return
	}

	insertID, err := h.userService.CreatUser(user)
	if err != nil {
		fmt.Println("", err.Error())
		responses.Error(w, 400, "could not creat user")
		return
	}
	responses.Success(w, insertID)
}
