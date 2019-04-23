package _routing

import (
	"errors"
	"github.com/gorilla/mux"
	"net/http"
)

type (
	FnsHandler func(http.ResponseWriter, *http.Request)

	Apihandler struct {
		FnsHandler
		Method  string
		Queries string
	}

	Handlers map[string]*Apihandler
)

func MainController(hs Handlers) (*mux.Router, error) {
	r := &mux.Router{}

	if len(hs) == 0 {
		return nil, errors.New("no handlers")
	}

	for p, h := range hs {
		r.HandleFunc(p, h.FnsHandler).Methods(h.Method)

		if h.Queries != "" {
			r.Queries(h.Queries)
		}
	}

	return r, nil
}
