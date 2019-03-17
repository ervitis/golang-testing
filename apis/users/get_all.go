package users

import (
	"encoding/json"
	"github.com/ervitis/golang-testing/helpers"
	"net/http"
	"strconv"
)

func (h *ReqHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []*User
	const itemsPerPage = 15
	var page int

	b, err := h.Reader.ReadData("data/users.json")
	if err != nil {
		helpers.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	_ = json.Unmarshal(b, &users)

	qpage := r.FormValue("page")
	if qpage == "" {
		page = 1
	} else {
		if page, err = strconv.Atoi(qpage); err != nil {
			helpers.ResponseWithError(w, http.StatusBadRequest, "page is not a number")
			return
		}
	}

	start := (page - 1) * itemsPerPage
	stop := start + itemsPerPage

	if start > len(users) {
		helpers.Response(w, http.StatusOK, []*User{})
		return
	}

	if stop > len(users) {
			stop = len(users)
	}


	helpers.Response(w, http.StatusOK, users[start:stop])
}
