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

	router.GET("/users", handlers.ListUsers)
	router.POST("/users", handlers.CreateUser)
	router.GET("/users/:user_id", handlers.IndexUser)
	router.PATCH("/users/:user_id", handlers.UpdateUser)
	router.DELETE("/users/:user_id", handlers.RemoveUser)

	log.Fatal(http.ListenAndServe(":8080", router))
}
