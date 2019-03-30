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

func (h *ReqHandler) GetCompany(w http.ResponseWriter, r *http.Request) {
	if r == nil {
		helpers.ResponseWithError(w, http.StatusBadRequest, "missing request")
		return
	}

	companyId := mux.Vars(r)["companyId"]
	if companyId == "" {
		helpers.ResponseWithError(w, http.StatusBadRequest, "missing companyId")
		return
	}

	b, err := h.Reader.ReadData("data/companies.json")
	if err != nil {
		helpers.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	uid, _ := strconv.Atoi(companyId)
	companyData := gojsonq.New().Reader(bytes.NewReader(b)).Where("id", "eq", uid).First()
	if companyData == nil {
		helpers.Response(w, http.StatusNotFound, companyData)
		return
	}

	c := &company{}
	b, _ = json.Marshal(companyData)
	_ = json.Unmarshal(b, c)

	helpers.Response(w, http.StatusOK, c)
}