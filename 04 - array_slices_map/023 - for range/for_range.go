package main

import "fmt"

func main() {

	numeros := [...]int{1, 2, 3, 4, 5} // Compilador Conta a Quantidade de Elementos

	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i, numero)
	}

	// Ignorando o Ìndice (localização do número dentro do Array)
	for _, num := range numeros {
		fmt.Println(num)
	}
}
