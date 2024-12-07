package app

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jaehanbyun/api/model"
	"github.com/urfave/negroni"
)

type AppHandler struct {
	http.Handler
	db model.DBHandler
}

func MakeHandler() *AppHandler {
	r := mux.NewRouter()
	neg := negroni.Classic()
	neg.UseHandler(r)

	a := &AppHandler{
		Handler: neg,
		db:      model.NewDBHandler(),
	}

	err := a.db.Init()
	if err != nil {
		panic(err)
	}

	r.HandleFunc("/health", a.healthCheck).Methods("GET")
	r.HandleFunc("/user", a.getUsers).Methods("GET")
	r.HandleFunc("/user/{id}", a.getUserByID).Methods("GET")
	r.HandleFunc("/user", a.createUser).Methods("POST")
	r.HandleFunc("/user/{id}", a.updateUser).Methods("PUT")
	r.HandleFunc("/user/{id}", a.deleteUser).Methods("DELETE")

	return a
}

func (a *AppHandler) healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte("OK")); err != nil {
		log.Printf("failed to write response: %v", err)
	}
}

func (a *AppHandler) getUsers(w http.ResponseWriter, r *http.Request) {
	users, err := a.db.GetUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(users); err != nil {
		log.Printf("failed to encode users: %v", err)
	}
}

func (a *AppHandler) getUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := a.db.GetUserByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if err := json.NewEncoder(w).Encode(user); err != nil {
		log.Printf("failed to encode user: %v", err)
	}
}

func (a *AppHandler) createUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = a.db.CreateUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (a *AppHandler) updateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var user model.User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user.ID = id
	err = a.db.UpdateUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (a *AppHandler) deleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = a.db.DeleteUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
