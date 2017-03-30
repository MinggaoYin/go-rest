package handlers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Welcome(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("Hello world!\n"))
}