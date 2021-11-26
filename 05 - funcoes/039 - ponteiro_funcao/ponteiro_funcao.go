package main

import "fmt"

func incremento1(n int) {
	n++ // -> n = n + 1
}

// REVISÃO: Um ponteiro é representado por um * (asterisco)
func incremento2(n *int) {
	/*
		REVISÃO: * é usado para desreferenciar, ou seja, ter acesso ao valor no qual o ponteiro aponta
	*/
	*n++
}

func main() {
	n := 1

	incremento1(n) // Opera sobre uma CÓPIA do valor da variável
	fmt.Println(n)

	// REVISÃO: & usado para obter o endereço da variável
	incremento2(&n) // Opera por Referência
	fmt.Println(n)
}
