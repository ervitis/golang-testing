package apis

import (
	"github.com/ervitis/golang-testing/apis/users"
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
	jr := &helpers.JsonReader{}

	apig := &users.GetHandler{
		Reader: jr,
	}

	return Handlers{
		"/users": apig.GetAllUsers,
	}
}
