package main

import (
	"cmd/handlers"
	"log"
	"net/http"
)

func main() {
	router := handlers.NewRouter(&handlers.Routes)
	log.Fatal(http.ListenAndServe(":8000", router))
}
