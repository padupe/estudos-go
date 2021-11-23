package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	// Pacote rand = Randon
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	// Limitamos o número aleatório que será gerado até o número 10
	return r.Intn(10)
}

func main() {
	// IF COM TRECHO DE INICIALIZAÇÃO
	if i := numeroAleatorio(); i > 5 { // Também é suportado no Switch
		fmt.Println("Ganhou!")
	} else {
		fmt.Println("Perdeu!")
	}
}
