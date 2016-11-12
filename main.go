package main

import (
	"log"
	"net/http"
)

func main() {

	httpRouter := CreateHttpRouter()

	log.Fatal(http.ListenAndServe(":8080", httpRouter))
}
