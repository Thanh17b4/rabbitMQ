package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"Thanh17b4/practice/handler/responses"
	"github.com/go-chi/chi/v5"
)

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
