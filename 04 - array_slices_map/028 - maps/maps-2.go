package main

import "fmt"

func main() {

	// Inicializando e declarando os valores do Map
	funcSalarios := map[string]float64{
		"Samuel Machado": 19000.50,
		"Rodrigo Prior":  50000.05,
		"Washington":     3987.75,
	}

	// Adicionando um novo elemento no Map
	funcSalarios["Rafael Filho"] = 1457.26

	// Se tentarmos deletar um elemento que não existe, teremos mensagem de erro
	// delete(funcSalarios, "Inexistente")

	for nome, salario := range funcSalarios {
		fmt.Println(nome, salario)
	}

	// OUTRA MANEIRA DE IMPRIMIR
	for nome, salario := range funcSalarios {
		fmt.Printf("%s (Salário: %f)", nome, salario)
	}
}
