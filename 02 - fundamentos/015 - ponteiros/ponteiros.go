package main

import "fmt"

func main() {
	i := 1

	// Go não tem aritmética de Ponteiros
	// Ponteiro é uma referência de Memória

	var p *int = nil // nil == null

	// & -> pegue o endereço da variável e me informe qual é
	p = &i

	// Peguei o valor associado ao Ponteiro
	*p++
	i++

	// Isso não é permitido
	// p++

	fmt.Println(p, *p, i, &i)
}
