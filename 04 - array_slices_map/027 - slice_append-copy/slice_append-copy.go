package main

import "fmt"

func main() {

	// APPEND -> Adiciona valores

	array1 := [3]int{7, 8, 9}
	var slice1 []int

	// O Código abaixo NÃO compila
	// array1 = append(array1, 4, 5, 6)

	slice1 = append(slice1, 4, 5, 6)
	fmt.Println(array1, slice1)

	slice2 := make([]int, 2)
	// Copiando do 'slice1' para o 'slice2'
	copy(slice2, slice1)
	fmt.Println(slice2)

	// IMPORTANTE: Só posso utilizar o COPY para Slice ou String

}
