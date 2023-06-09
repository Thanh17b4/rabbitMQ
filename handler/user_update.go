package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"Thanh17b4/practice/handler/responses"
	"Thanh17b4/practice/model"
	"github.com/go-chi/chi/v5"
)

func (h UserHandle) UpdateUserHandle(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		responses.Error(w, r, 400, err, "userID must be number")
		return
	}
	reqBody, _ := ioutil.ReadAll(r.Body)
	var user *model.User
	err1 := json.Unmarshal(reqBody, &user)
	if err1 != nil {
		responses.Error(w, r, 400, err1, "could not Unmarshal body request")
		return
	}
	user.ID = int(id)
	updatedUser, err := h.userService.UpdateUserService(user)
	if err != nil {
		responses.Error(w, r, 400, err, "Could not update user because information input is not correct")
		return
	}
	responses.Success(w, r, http.StatusAccepted, updatedUser)
}
