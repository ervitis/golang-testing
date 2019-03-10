package users

import (
	"encoding/json"
	"github.com/ervitis/golang-testing/helpers"
	"net/http"
)

func (h *ReqHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []*User

	b, err := h.Reader.ReadData("data/users.json")
	if err != nil {
		helpers.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	_ = json.Unmarshal(b, &users)

	helpers.Response(w, http.StatusOK, users)
}
