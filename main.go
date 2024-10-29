package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	ID    int
	Nome  string
	Email string
}

type UserDTO struct {
	Nome  string
	Email string
}

var users = []User{
	{
		ID:    1,
		Nome:  "Ana",
		Email: "ana@teste.com",
	},
	{
		ID:    2,
		Nome:  "Geovana",
		Email: "geovana@teste.com",
	},
}

func obteruser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func addUser(w http.ResponseWriter, r *http.Request) {
	var user User
	var dto UserDTO
	json.NewDecoder(r.Body).Decode(&dto)

	user.ID = len(users) + 1
	user.Email = dto.Email
	user.Nome = dto.Nome

	users = append(users, user)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func main() {
	fmt.Println("API")

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			obteruser(w, r)
		} else if r.Method == http.MethodPost {
			addUser(w, r)
		} else {
			http.Error(w, "Methos invalido", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Servidor rodando na porta: 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
