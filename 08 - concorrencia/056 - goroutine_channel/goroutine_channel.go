package main

import (
	"fmt"
	"time"
)

// Channel (canal): é a forma de comunicação entre as GoRoutines
// Channel é um TIPO (como float, int, func...)

func doisTresQuatroVezes(base int, c chan int) {
	// Tempo de "espera" para executar de 1 segundo
	time.Sleep(time.Second)
	c <- 2 * base // Estou enviando dados para o Channel

	// Tempo de "espera" para executar de 1 segundo
	time.Sleep(time.Second)
	c <- 3 * base // Estou enviando dados para o Channel

	// Tempo de "espera" para executar de 1 segundo
	time.Sleep(3 * time.Second)
	c <- 4 * base // Estou enviando dados par ao Channel
}

func main() {
	c := make(chan int)
	go doisTresQuatroVezes(2, c) // GoRoutine simula um assíncronismo...
	fmt.Println("Passo 1")

	a, b := <-c, <-c // Estou recebendo os dados do Channel
	fmt.Println("Passo 2")
	fmt.Println(a, b)

	fmt.Println("Passo 3")
	fmt.Println(<-c)
}
