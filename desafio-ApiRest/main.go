package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Client struct {
	ID       int
	Nome     string `json:"nome"`
	Email    string `json:"email"`
	Telefone string `json:"telefone"`
	Endereco string `json:"endereco"`
}

var Clients []Client
var proximoId = 1

func HandleAddPerson(w http.ResponseWriter, r *http.Request) {
	var p Client

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if p.Nome == "" || p.Email == "" {
		http.Error(w, "Nome e e-mail são obrigatórios", http.StatusBadRequest)
		return
	}

	p.ID = proximoId
	proximoId++
	Clients = append(Clients, p)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(p)
}

func HandleHello(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if nomeDoJogador := vars["name"]; nomeDoJogador != "" {
		fmt.Fprintf(w, "Hello, %s\n", nomeDoJogador)
	} else {
		fmt.Fprint(w, "Hello, World!\n")
	}
}

func HandleHelloJson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": fmt.Sprintf("Hello, %s", vars["name"]),
	})
}

func mostrarContatos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Clients)
}

func main() {
	r := mux.NewRouter()

	// r.HandleFunc("/", HandleHello).Methods("GET")
	// r.HandleFunc("/{name}", HandleHello).Methods("GET")

	// r.HandleFunc("/json/{name}", HandleHelloJson).Methods("GET")
	r.HandleFunc("/dados", mostrarContatos).Methods("GET")

	r.HandleFunc("/person", HandleAddPerson).Methods("POST")

	http.ListenAndServe(":8000", r)
}
