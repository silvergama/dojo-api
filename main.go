package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/silvergama/dojo-api/entity"
	"github.com/silvergama/dojo-api/repository"
)

type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Age       int    `json:"age"`
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("OK"))
	if err != nil {
		fmt.Println("Error")
	}
}

func allUsers(w http.ResponseWriter, r *http.Request) {
	users, err := repository.GetUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func addUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Add("content-type", "application/json")

	user := &User{}
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	u, err := entity.NewUser(user.FirstName, user.LastName, user.Email, user.Phone)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if err = repository.AddUser(*u); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		allUsers(w, r)
	case http.MethodPost:
		addUser(w, r)
	default:
		http.Error(w, "Invalid requerest method", http.StatusMethodNotAllowed)
	}
}

func main() {
	err := repository.Setup()
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/healthcheck", HealthCheck)
	http.HandleFunc("/users", UserHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
