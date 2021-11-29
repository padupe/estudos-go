package main

import (
	"fmt"
	"time"
)

// Channel com Buffer: Permite que se criem posições, que só quando "cheias" geram um bloqueio

func rotina(canal chan int) {
	fmt.Println("Início da Execução")

	canal <- 1
	fmt.Println("Executou o 1")
	canal <- 2
	fmt.Println("Executou o 2")
	canal <- 3
	fmt.Println("Executou o 3")
	canal <- 4
	fmt.Println("Será que executa o 4?")
	canal <- 5
	canal <- 6
}

func main() {
	// Canal que espera INT, com Buffer de 3
	canal := make(chan int, 3)
	go rotina(canal)

	time.Sleep(time.Second)
	fmt.Println(<-canal)
	// Não executa o 4, pois o Buffer limitou até 3 elementos
	// O 4 só seria executado, caso um dos elementos anteriores fosse consumido
}
