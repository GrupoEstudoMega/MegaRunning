package main

import (
	"./prova"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/api/provas", prova.Handler) //.Methods("GET")
	// r.HandleFunc("/api/provas", CreateKittenHandler).Methods("POST")
	r.HandleFunc("/api/provas/{id}", prova.HandlerId) //.Methods("DELETE")
	http.Handle("/api/", r)

	// http.HandleFunc("/api/provas")
	fs := http.Dir("public/")
	http.Handle("/", http.FileServer(fs))
	fmt.Println("Rodando na porta 8080")
	http.ListenAndServe(":8080", nil)
}
