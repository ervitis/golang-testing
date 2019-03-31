package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/ervitis/golang-testing/helpers"
	"github.com/thedevsaddam/gojsonq"
	"net/http"
	"strconv"
)

func (h *ReqHandler) GetAllCompanies(w http.ResponseWriter, r *http.Request) {
	var b []byte
	var st int

	b, err := h.Reader.ReadData("data/companies.json")
	if err != nil {
		helpers.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	qpage := r.URL.Query().Get("page")

	quser := r.URL.Query().Get("userId")
	if quser != "" {
		if b, st = getCompaniesByUser(b, quser); st != http.StatusOK {
			helpers.ResponseWithError(w, st, "error getting companies by user")
			return
		}
	}

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

func getCompaniesByUser(data []byte, userId string) ([]byte, int) {
	uid, _ := strconv.Atoi(userId)

	companyData := gojsonq.New().Reader(bytes.NewReader(data)).Where("userId", "eq", uid).Get()

	if v, ok := companyData.([]interface{}); !ok {
		return nil, http.StatusInternalServerError
	} else {
		if len(v) == 0 {
			return nil, http.StatusNotFound
		}
	}

	b, _ := json.Marshal(companyData)

	return b, http.StatusOK
}
