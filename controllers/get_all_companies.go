package controllers

import (
	"encoding/json"
	"github.com/ervitis/golang-testing/helpers"
	"net/http"
)

func (h *ReqHandler) GetAllCompanies(w http.ResponseWriter, r *http.Request) {

	b, err := h.Reader.ReadData("data/companies.json")
	if err != nil {
		helpers.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	qpage := r.URL.Query().Get("page")

	var companies []*company
	if data, err := helpers.NewPaginator(b).Paginate(qpage); err != nil {
		helpers.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	} else {
		b, _ = json.Marshal(data)
		json.Unmarshal(b, &companies)
	}

	helpers.Response(w, http.StatusOK, companies)
}
