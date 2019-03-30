package main

import (
	"github.com/ervitis/golang-testing/routing"
	"github.com/ervitis/golang-testing/server"
	"log"
	"net/http"
)

func main() {
	srv := &server.Server{Addr: "http://localhost", Port: "8080"}
	log.Fatal(http.ListenAndServe(srv.PortAddress(), routing.MainController(routing.Routes())))
}
