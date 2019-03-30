package routing

import (
	"github.com/ervitis/golang-testing/controllers"
	"github.com/ervitis/golang-testing/helpers"
	"github.com/gorilla/mux"
	"net/http"
)

type fnsHandler func(http.ResponseWriter, *http.Request)
type Handlers map[string]fnsHandler

func MainController(hs Handlers) *mux.Router {
	r := &mux.Router{}

	for p, h := range hs {
		r.HandleFunc(p, h)
	}

	return r
}

func Routes() Handlers {
	apis := &controllers.ReqHandler{Reader: &helpers.JsonReader{}}

	return Handlers{
		"/users":         apis.GetAllUsers,
		"/user/{userId}": apis.GetUser,
		"/companies":     apis.GetAllCompanies,
	}
}
