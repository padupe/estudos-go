package main

import (
	"fmt"
	"strings"
)

type pessoa struct {
	nome      string
	sobrenome string
}

func (p pessoa) getNomeCompleto() string {
	return p.nome + " " + p.sobrenome
}

func (p *pessoa) setNomeCompleto(nomeCompleto string) {
	partes := strings.Split(nomeCompleto, " ")
	p.nome = partes[0]
	p.sobrenome = partes[1]
}

func main() {
	pessoa1 := pessoa{"Pedro", "Silva"}
	fmt.Println(pessoa1.getNomeCompleto())

	pessoa2 := pessoa{nome: "João", sobrenome: "Santos"}
	fmt.Println(pessoa2.getNomeCompleto())

	// Utilizando-se do Ponteiro
	pessoa1.setNomeCompleto("Maria Pereira")
	fmt.Println(pessoa1.getNomeCompleto())
}
