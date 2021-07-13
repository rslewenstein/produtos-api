package main

import (
	"fmt"
	"log"
	"net/http"

	"produtos-api/controller"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Alive...")
}

func handleRequests() {
	myRoute := mux.NewRouter().StrictSlash(true)
	myRoute.HandleFunc("/", home)
	myRoute.HandleFunc("/listar", controller.ListarProdutos()).Methods("GET")
	myRoute.HandleFunc("/adicionar", controller.AdicionarProdutos()).Methods("POST")
	log.Fatal(http.ListenAndServe(":5000", myRoute))
}

func main() {
	handleRequests()
}
