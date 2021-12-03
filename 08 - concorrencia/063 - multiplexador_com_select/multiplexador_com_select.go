package main

import (
	"fmt"
	"time"
)

func falar(pessoa string) <-chan string {

	channel := make(chan string)

	// Funções Anônimas que não recebem parâmetros, devem ser encerradas com ()
	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(time.Second)
			// Estou enviando dados para o Channel
			channel <- fmt.Sprintf("%s falando: %d", pessoa, i)
		}
	}()

	return channel
}

func juntar(entrada1, entrada2 <-chan string) <-chan string {

	channel2 := make(chan string)
	go func() {
		// FOR Infinito
		for {
			select {
			case s := <-entrada1:
				channel2 <- s
			case s := <-entrada2:
				channel2 <- s
			}
		}
	}()

	return channel2
}

func main() {
	channel3 := juntar(falar("João"), falar("Maria"))
	fmt.Println(<-channel3, <-channel3)
	fmt.Println(<-channel3, <-channel3)
	fmt.Println(<-channel3, <-channel3)
}

//git commit -m $'feat(08 - concorrencia/063 - multiplexador_com_select): Adicionado ao Projeto.\nAula 83 - Curso Go (Golang): Explorando a linguagem do Google.\nModulo 08 Finalizado.'
