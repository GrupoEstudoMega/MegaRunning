package main

import (
	"./prova"
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/prova", prova.Handler)
	fmt.Println("Rodando na porta 8080")
	http.ListenAndServe(":8080", nil)

}

