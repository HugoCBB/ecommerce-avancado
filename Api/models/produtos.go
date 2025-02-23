package models

import "time"

type Produtos struct {
	ID          int       `json:"id"`
	Nome        string    `json:"nome"`
	Descricao   string    `json:"descricao"`
	Preco       int       `json:"preco"`
	DataCriacao time.Time `json:"data_criacao"`
}

var Produto []Produtos
