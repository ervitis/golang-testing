package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/ervitis/golang-testing/helpers"
	"github.com/gorilla/mux"
	"github.com/thedevsaddam/gojsonq"
	"net/http"
	"strconv"
)

func (h *ReqHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	if r == nil {
		helpers.ResponseWithError(w, http.StatusBadRequest, "missing request")
		return
	}

	userId := mux.Vars(r)["userId"]
	if userId == "" {
		helpers.ResponseWithError(w, http.StatusBadRequest, "missing userId")
		return
	}

	b, err := h.Reader.ReadData("data/users.json")
	if err != nil {
		helpers.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	uid, _ := strconv.Atoi(userId)
	userData := gojsonq.New().Reader(bytes.NewReader(b)).Where("id", "eq", uid).First()
	if userData == nil {
		helpers.Response(w, http.StatusNotFound, userData)
		return
	}

	u := &user{}
	b, _ = json.Marshal(userData)
	_ = json.Unmarshal(b, u)

	helpers.Response(w, http.StatusOK, u)
}
