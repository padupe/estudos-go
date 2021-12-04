package main

import (
	"fmt"
	"strconv"
	// c "stringconv" -> para usar alias nos pacotes, ou como no JS, declarar uma variável para utilizar o pacote
)

func main() {
	x := 2.4 // tipo float64
	y := 2

	fmt.Println(x / float64(y)) // Preciso converter o valor de y para float64
	// Retorno esperado: 1.2

	nota := 6.9
	notaFinal := int(nota) // Arredonda para menos =(
	fmt.Println(notaFinal)
	// Retorno esperado: 6

	// CUIDADO AO CONCATENAR STRINGS COM INTEIROS
	fmt.Println("Teste " + string(97))
	// Resultado obtido será: Teste a -> O código compila o caracter referente na tabela Unicode

	// INTEIRO PARA STRING
	fmt.Println("Teste " + strconv.Itoa(97))
	// Retorno esperado: Teste 97

	// STRING PARA INTEIRO
	num, _ := strconv.Atoi("123") // Primeiro parâmetro é o retorno, segundo parâmetro um erro (neste caso, não estamos tratando o erro)
	fmt.Println(num - 122)        // 123 - 122
	// Retorno esperado: 1

	b, _ := strconv.ParseBool("true") // Primeiro parâmetro é o retorno, segundo parâmetro um erro (neste caso, não estamos tratando o erro)
	if b {
		fmt.Println("Verdadeiro")
	}
	// Resultado esperado: Verdadeiro
}
