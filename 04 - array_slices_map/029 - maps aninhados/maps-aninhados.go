package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Geraldo": 15456.78,
			"Gustavo": 10458.15,
		},
		"J": {
			"João":    4587.17,
			"Juliana": 3000.00,
		},
		"M": {
			"Maria": 1987.15,
		},
		"P": {
			"Pedro": 2457.70,
		},
	}

	fmt.Println(funcsPorLetra)

	delete(funcsPorLetra, "P")
	fmt.Println(funcsPorLetra)

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)

		// Desafio Proposto na aula, realizado de maneira individual e autônoma
		for nome, salario := range funcs {
			fmt.Println(nome, salario)
		}
	}
}
