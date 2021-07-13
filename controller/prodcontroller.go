package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"produtos-api/model"
)

func ListarProdutos() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prod := model.Produtos{
			model.Produto{Id: 1, Nome: "Pó de Café", Fabricante: "Cocô de Jacú", Lote: "As111"},
		}
		json.NewEncoder(w).Encode(prod)
	}
}

func AdicionarProdutos() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// prod := model.Produtos{
		// 	model.Produto{Id: 1, Nome: "Pó de Café", Fabricante: "Cocô de Jacú", Lote: "As111"},
		// }
		// json.NewEncoder(w).Encode(prod)

		fmt.Fprintf(w, "Teste de POST")
	}
}
