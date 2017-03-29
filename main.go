package main

import (
	"log"
	"net/http"

	"go-rest/handlers"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	router.GET("/", handlers.Welcome)

	log.Fatal(http.ListenAndServe(":8080", router))
}
