package main

import (
	"log"
	"net/http"

	"go-rest/handlers"
)

func main() {
	http.HandleFunc("/", handlers.Welcome)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
