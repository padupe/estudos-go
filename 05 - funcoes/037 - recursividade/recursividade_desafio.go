package main

import "fmt"

// RECURSIVIDADE -> A função chama ela mesma

func fatorial(n uint) uint {

	switch {
	case n == 0:
		return 1
	default:
		return n * fatorial(n-1)
	}
}

func main() {

	resultado := fatorial(5)
	fmt.Println(resultado)

	// O código entre as linhas 23 e 24 não é possível, pois validei apenas números SEM SINAL -> uint
	// valorNegativo := fatorial(-4)
	// fmt.Println(valorNegativo)
}
