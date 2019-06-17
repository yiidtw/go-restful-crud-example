package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// User is ...(go-lint is expected some comments here)
type User struct {
	ID       int    `json:"id"`
	UserName string `json:"username"`
}

var users []User

// Get All Users
func getAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newUser User
	json.NewDecoder(r.Body).Decode(&newUser)

	if len(newUser.UserName) > 0 {
		newUser.ID = len(users) + 1
		users = append(users, newUser)
		json.NewEncoder(w).Encode(newUser)
	}
}

func getSingleUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// get ID
	params := mux.Vars(r)
	for _, u := range users {
		if strconv.Itoa(u.ID) == params["id"] {
			json.NewEncoder(w).Encode(u)
			return
		}
	}
}

func updateSingleUserName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newUser User
	json.NewDecoder(r.Body).Decode(&newUser)
	fmt.Print(newUser.ID)
	fmt.Print(newUser.UserName)

	params := mux.Vars(r)
	for i, u := range users {
		if strconv.Itoa(u.ID) == params["id"] {
			users = append(users[:i], users[i+1:]...)

			u.UserName = newUser.UserName
			users = append(users, u)
			json.NewEncoder(w).Encode(u)
			return
		}
	}
}

func deleteSingleUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, u := range users {
		if strconv.Itoa(u.ID) == params["id"] {
			users = append(users[:i], users[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(users)
}

func main() {
	users = append(users,
		User{ID: 1, UserName: "Frank"},
		User{ID: 2, UserName: "Lisa"},
		User{ID: 3, UserName: "Penny"},
	)
	router := mux.NewRouter()
	router.HandleFunc("/users", getAllUsers).Methods("GET")
	router.HandleFunc("/user", createUser).Methods("POST")
	router.HandleFunc("/user/{id}", getSingleUser).Methods("GET")
	router.HandleFunc("/user/{id}", updateSingleUserName).Methods("POST")
	router.HandleFunc("/user/{id}", deleteSingleUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
