package controllers

import (
	"github.com/ervitis/golang-testing/controllers/_routing"
	"github.com/ervitis/golang-testing/helpers"
	"net/http"
)

func Routes() _routing.Handlers {
	apis := &ReqHandler{Reader: &helpers.JsonReader{}}

	return _routing.Handlers{
		"/users":               &_routing.Apihandler{FnsHandler: apis.GetAllUsers, Method: http.MethodGet, Queries: "page"},
		"/user/{userId}":       &_routing.Apihandler{FnsHandler: apis.GetUser, Method: http.MethodGet},
		"/companies":           &_routing.Apihandler{FnsHandler: apis.GetAllCompanies, Method: http.MethodGet},
		"/company/{companyId}": &_routing.Apihandler{FnsHandler: apis.GetCompany, Method: http.MethodGet},
	}
}
