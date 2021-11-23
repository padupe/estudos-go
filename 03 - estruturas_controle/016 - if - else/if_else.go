package main

import "fmt"

func imprimirResultado(nota float64) {
	// O parâmetro não precisa ser informado entre parênteses
	if nota >= 7 {
		fmt.Println("Aprovado com nota", nota)
	} else {
		fmt.Println("Reprovado com nota", nota)
	}
}

func main() {
	imprimirResultado(7.3)
	imprimirResultado(5.5)
}
