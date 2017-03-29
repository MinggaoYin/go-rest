package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", Welcome)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!\n"))
}