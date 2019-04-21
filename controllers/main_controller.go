package controllers

import (
	"errors"
	"github.com/gorilla/mux"
	"net/http"
)

type (
	fnsHandler func(http.ResponseWriter, *http.Request)

	apihandler struct {
		fnsHandler
		method  string
		queries string
	}

	Handlers map[string]*apihandler
)

func MainController(hs Handlers) (*mux.Router, error) {
	r := &mux.Router{}

	if len(hs) == 0 {
		return nil, errors.New("no handlers")
	}

	for p, h := range hs {
		r.HandleFunc(p, h.fnsHandler).Methods(h.method)

		if h.queries != "" {
			r.Queries(h.queries)
		}
	}

	return r, nil
}
