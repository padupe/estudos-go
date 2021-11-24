package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3} // Array

	s1 := []int{1, 2, 3} // Slice [Parte]

	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}

	// IMPORTANTE: Slice não é um Array! Slice define o 'pedaço' de um Array

	// No exemplo abaixo, o Slice inclui o elemento na posição 1 até 3, porém não inclui a posicão 3
	// Em resumo, o Slice captura a posição 1 e 2
	s2 := a2[1:3]
	fmt.Println(a2, s2)

	// No exemplo abaixo, o Slice inclui o elemento na posição 0 até 2, porém não inclui a posicão 2
	// Em resumo, o Slice captura a posição 0 e 1
	s3 := a2[:2] // Novo Slice, mas considerando o MESMO Array -> a2
	fmt.Println(a2, s3)

	// Podemos considerar Slice como: TAMANHO e um PONTEIRO

	/*
		SLICE DO SLICE
		No exemplo abaixo, o Slice inclui o elemento na posição 0 até 1, porém não inclui a posicão 1
		Em resumo, o Slice (do Slice) captura a posição 0

		Retorno previsto: 2
	*/
	s4 := s2[:1]
	fmt.Println(s2, s4)
}
