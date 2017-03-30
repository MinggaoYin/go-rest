package handlers

import (
	"encoding/json"
	"net/http"

	"go-rest/resources"

	"github.com/julienschmidt/httprouter"
)

func Welcome(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("Hello world!\n"))
}

func IndexUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userId := ps.ByName("user_id")
	user, err := resources.GetUser(userId)
	if err == resources.ErrNotFound {
		w.WriteHeader(404)
		return
	}
	b, _ := json.Marshal(user)
	w.Write(b)
}

func CreateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var params *resources.User
	json.NewDecoder(r.Body).Decode(&params)
	b, _ := json.Marshal(resources.AddUser(params.Name, params.Age, params.Email))
	w.Write(b)
}

func UpdateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var params map[string]interface{}
	json.NewDecoder(r.Body).Decode(&params)

	userId := ps.ByName("user_id")
	updatedUser, err := resources.UpdateUser(userId, params)
	if err == resources.ErrNotFound {
		w.WriteHeader(404)
		return
	}

	b, _ := json.Marshal(updatedUser)
	w.Write(b)
}

func RemoveUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userId := ps.ByName("user_id")
	resources.RemoveUser(userId)
	w.WriteHeader(204)
}

func ListUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	b, _ := json.Marshal(resources.ListUsers())
	w.Write(b)
}
