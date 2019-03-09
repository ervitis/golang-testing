package users

import (
	"github.com/ervitis/golang-testing/helpers"
	"net/http"
)

type GetHandler struct {
	Reader helpers.Reader
}

func (h *GetHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []*User

	if err := h.Reader.ReadData("data/users.json", &users); err != nil {
		helpers.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.Response(w, http.StatusOK, users)
	return
}
