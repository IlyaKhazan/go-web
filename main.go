package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/jackc/pgx/v5"
	"log"
	"net/http"
)

type User struct {
	Firstname string
	Lastname  string
	Email     string
	Age       uint
}

var users []User

func main() {
	http.HandleFunc("/users", usersHandler)

	log.Println("Listening on port 8080")
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func usersHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		getUsers(w, r)
	case http.MethodPost:
		postUsers(w, r)
	default:
		http.Error(w, "Invalid http method", http.StatusMethodNotAllowed)
	}
}

func getUsers(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(users)

}

func postUsers(w http.ResponseWriter, r *http.Request) {

	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	users = append(users, user)
	fmt.Fprintf(w, "Пользователь добавлен: '%v'\n", user)
}
