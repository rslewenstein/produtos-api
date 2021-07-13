package model

type Produto struct {
	Id         int    `json:"id"`
	Nome       string `json:"nome"`
	Fabricante string `json:"fabricante"`
	Lote       string `json:"lote"`
}

type Produtos []Produto
