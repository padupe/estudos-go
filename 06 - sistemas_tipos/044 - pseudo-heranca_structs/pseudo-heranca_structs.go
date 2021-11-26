package main

import "fmt"

// ATENÇÃO: GO NAÕ POSSUI HERANÇA

type carro struct {
	nome            string
	velocidadeAtual int
}

type ferrari struct {
	carro       // O Go cria campos Anônimos
	turboLigado bool
}

func main() {
	f := ferrari{}
	f.nome = "F40"
	f.velocidadeAtual = 0
	f.turboLigado = true

	// %s refere-se a STRING (f.nome) | %t refere-se a BOOLEAN (f.turboLigado)
	fmt.Printf("A ferrari %s está com turbo ligado?\n%t\n", f.nome, f.turboLigado)

	// Imprimindo a Estrutura de 'f'
	fmt.Println(f)
}
