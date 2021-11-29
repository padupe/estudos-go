package main

import "fmt"

func main() {
	// Abaixo estamos criando um canal (channel) de inteiros | segundo parâmetro é um Buffer [Será abordado no decorrer do curso]
	canal := make(chan int, 1)

	// Estou enviando para o canal o valor '1', ou seja, este é o processo de Escrita
	canal <- 1

	// Estou recebendo dados do canal, ou seja, este é o processo de Leitura
	<-canal

	canal <- 2
	fmt.Println(<-canal)
}
