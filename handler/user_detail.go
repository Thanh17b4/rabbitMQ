package handler

import (
	"net/http"
	"strconv"

	"Thanh17b4/practice/handler/responses"
	"github.com/go-chi/chi/v5"
)

func (h UserHandle) GetDetailUserHandle(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "id")

	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		responses.Error(w, r, 400, err, "userID must be number")
		return
	}
	user, err := h.userService.GetDetailUser(id)
	if err != nil {
		responses.Error(w, r, 400, err, "could not get user")
		return
	}
	responses.Success(w, r, http.StatusAccepted, user)
}
