package main

import (
	"fmt"
	"log"
	"net/http"

	"produtos-api/controller"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Alive...")
}

func handleRequests() {
	http.HandleFunc("/", home)
	http.HandleFunc("/listar", controller.ListarProdutos())
	log.Fatal(http.ListenAndServe(":5000", nil))
}

func main() {
	handleRequests()
}
