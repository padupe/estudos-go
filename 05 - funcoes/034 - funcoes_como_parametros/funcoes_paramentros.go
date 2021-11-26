package main

import "fmt"

// Funmção sendo informada como Parâmetro para OUTRA FUNÇÃO
func multiplicacao(a, b int) int {
	return a * b
}

func exec(funcao func(int, int) int, param1, param2 int) int {
	return funcao(param1, param2)
}

func main() {
	resultado := exec(multiplicacao, 3, 4)
	fmt.Println(resultado)
	// Resultado Esperado: 12
}
