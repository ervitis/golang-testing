package controllers

import (
	"github.com/ervitis/golang-testing/helpers"
	"net/http"
)

func Routes() Handlers {
	apis := &ReqHandler{Reader: &helpers.JsonReader{}}

	return Handlers{
		"/users":               &apihandler{fnsHandler: apis.GetAllUsers, method: http.MethodGet, queries: "page"},
		"/user/{userId}":       &apihandler{fnsHandler: apis.GetUser, method: http.MethodGet},
		"/companies":           &apihandler{fnsHandler: apis.GetAllCompanies, method: http.MethodGet},
		"/company/{companyId}": &apihandler{fnsHandler: apis.GetCompany, method: http.MethodGet},
	}
}
