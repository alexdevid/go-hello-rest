package main

import (
	"github.com/alexdevid/go-hello-rest/application"
	"log"
	"net/http"
)

func main() {

	router := application.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
