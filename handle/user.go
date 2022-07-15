package handle

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	model "practice/model"
	"strconv"
)

type UserService interface {
	GetListUser(page int64, limit int64) ([]*model.User, error)
	GetDetailUser(userID int64) (*model.User, error)
	UpdateUserService(user *model.User) *model.User
	DeleteUser(userID int64) (int64, error)
	CreatUser(user *model.User) (*model.User, error)
}
type UserHandle struct {
	userHandle UserService
}

func NewUserHandle(userHandle UserService) *UserHandle {
	return &UserHandle{userHandle: userHandle}
}
func (h UserHandle) GetListUser(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	currentPage, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		fmt.Println("page must be number: ", err.Error())
		return
	}
	fmt.Println(" CurrentPage: ", currentPage)
	perPage := r.URL.Query().Get("perPage")
	limit, err := strconv.ParseInt(perPage, 10, 64)
	if err != nil {
		fmt.Println("page must be number: ", err.Error())
		return
	}
	fmt.Println(" perPage: ", limit)
	users, err := h.userHandle.GetListUser(currentPage, limit)
	for _, user := range users {
		fmt.Println("user: ", user)
		json.NewEncoder(w).Encode(user)
	}
}

func (h UserHandle) GetDetailUserHandle(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("ddd")
	vars := mux.Vars(r)
	userID := vars["id"]
	//fmt.Println("kkk")
	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		fmt.Println(" ID must be number")
		return
	}
	//fmt.Println("eeee")
	user, err := h.userHandle.GetDetailUser(id)
	if err != nil {
		fmt.Println("id is not valid: ", err.Error())
		return
	}
	//fmt.Println("id: ", id)
	json.NewEncoder(w).Encode(user)
}

func (h UserHandle) UpdateUserHandle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var user *model.User
	err := json.Unmarshal(reqBody, &user)
	if err != nil {
		log.Fatalln("could not Unmarshal body request")
		return
	}
	vars := mux.Vars(r)
	userID := vars["id"]
	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		fmt.Printf("invalid id. ID should be number")
		return
	}
	user.ID = int(id)
	fmt.Println(user)
	updateID := h.userHandle.UpdateUserService(user)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"updatedID": updateID,
	})
}

func (h UserHandle) DeleteUserHandle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]
	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		fmt.Println(" id must be number: ", err.Error())
		return
	}
	deleteId, _ := h.userHandle.DeleteUser(id)
	json.NewEncoder(w).Encode(map[string]interface{}{
		" Deleted row ": deleteId,
	})
}

func (h UserHandle) CreatUserHandle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var user *model.User
	err := json.Unmarshal(reqBody, &user)
	if err != nil {
		fmt.Println(" can not marshal your request: ", err.Error())
		return
	}

	insertID, err := h.userHandle.CreatUser(user)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"register successfully": insertID,
	})
}
