package main

import (
	"github.com/ervitis/golang-testing/controllers"
	"github.com/ervitis/golang-testing/server"
	"net/http"
)

func main() {
	srv := &server.Server{Addr: "http://localhost", Port: "8080"}
	r, err := controllers.MainController(controllers.Routes())
	if err != nil {
		panic(err)
	}

	if err = http.ListenAndServe(srv.PortAddress(), r); err != nil {
		panic(err)
	}
}
