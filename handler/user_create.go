package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/Thanh17b4/practice/handler/responses"

	model "github.com/Thanh17b4/practice/model"
	"github.com/go-chi/chi/v5"
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

func (h UserHandle) GetDetailUserHandle(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "id")
	fmt.Println(userID)
	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		responses.Error(w, r, 400, err, userID)
		//responses.Error(w, r, 400, err, "userID must be number")
		return
	}
	user, err := h.userService.GetDetailUser(id)
	if err != nil {
		responses.Error(w, r, 400, err, "could not get user")
		return
	}
	responses.Success(w, r, http.StatusAccepted, user)
}

func (h UserHandle) GetListUser(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	currentPage, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		responses.Error(w, r, 400, err, "page must be number")
		return
	}
	perPage := r.URL.Query().Get("perPage")
	limit, err := strconv.ParseInt(perPage, 10, 64)
	if err != nil {
		responses.Error(w, r, 400, err, "perPage must be number")
		return
	}
	users, err := h.userService.GetListUser(currentPage, limit)
	for _, user := range users {
		responses.Success(w, r, http.StatusAccepted, user)
	}
}

func (h UserHandle) UpdateUserHandle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var user *model.User
	err := json.Unmarshal(reqBody, &user)
	if err != nil {
		responses.Error(w, r, 400, err, "could not Unmarshal body request")
		return
	}
	userID := chi.URLParam(r, "id")

	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		responses.Error(w, r, 400, err, "userID must be number")
		return
	}
	user.ID = int(id)
	updateID, err := h.userService.UpdateUserService(user)
	if err != nil {
		responses.Error(w, r, 400, err, "Could not update user because information input is not correct")
		return
	}
	responses.Success(w, r, http.StatusAccepted, updateID)
}

func (h UserHandle) DeleteUserHandle(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		responses.Error(w, r, 400, err, "userID must be number")
		return
	}
	deleteId, err := h.userService.DeleteUser(id)
	if err != nil {
		responses.Error(w, r, http.StatusBadRequest, err, "Could not delete user")
		return
	}
	responses.Success(w, r, http.StatusAccepted, fmt.Sprintf("userID %d has been deleted", deleteId))
}
