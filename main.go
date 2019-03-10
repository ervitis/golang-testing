package main

import (
	"github.com/ervitis/golang-testing/apis"
	"github.com/ervitis/golang-testing/server"
	"log"
	"net/http"
)

func main() {
	server := &server.Server{Addr: "http://localhost", Port: "8080"}
	log.Fatal(http.ListenAndServe(server.PortAddress(), apis.MainController(apis.Routes())))
}
