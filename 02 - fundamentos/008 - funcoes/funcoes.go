package main

import "fmt"

func somar(a int, b int) int { // (nome do parâmetro e tipo do parâmetro, nome do parâmetro e tipo do parâmetro) retorno
	return a + b
}

func imprimir(valor int) {
	fmt.Println(valor)
}

/*

ESTE TRECHO DE CÓDIGO (Linhas 17 a 20) FOI TRANSFERIDO PARA O ARQUIVO 'main.go'

func main() {
	resultado := somar(3, 4)
	imprimir(resultado)
}
-> Será impresso o valor 7
*/

/*
	COMO EXECUTAR ESTE CÓDIGO PELO TERMINAL
	cd 02. fundamentos/funcoes
	go run *.go
*/
