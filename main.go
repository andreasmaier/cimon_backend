package main

import (
	"net/http"
	"log"
)

func main() {
	go h.run()

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":3000", router))
}