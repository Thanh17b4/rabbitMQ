package handler

import (
	"net/http"
	"strconv"

	"github.com/Thanh17b4/practice/handler/responses"
)

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
	users, err1 := h.userService.GetListUser(currentPage, limit)
	if err1 != nil {
		responses.Error(w, r, 400, err1, "Could not get list users")
	}
	for _, user := range users {
		responses.Success(w, r, 200, user)
	}
}
