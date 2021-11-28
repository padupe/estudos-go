package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	// Identicador iniciado com LETRA MAIÚSCULA indica que ele será Privado
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	// Struct para JSON

	p1 := produto{1, "Notebook", 2459.99, []string{"Promoção", "Eletrônico", "Informatica"}}

	// estou ignorando o erro no segundo parâmetro
	p1JSON, _ := json.Marshal(p1)
	fmt.Println(string(p1JSON))

	// JSON para Struct
	var p2 produto
	jsonToSTRING := `{"id":2, "nome":"Caneta", "preco":8.90, "tags":["Papelaria","Importado"]}`
	json.Unmarshal([]byte(jsonToSTRING), &p2)
	fmt.Println(p2)
	fmt.Println(p2.Tags[1])
}
