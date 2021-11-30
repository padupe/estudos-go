package main

import (
	"fmt"

	html "github.com/padupe/titulos-go"
)

func encaminhar(origem <-chan string, destino chan string) {
	for {
		// Estou enviando dados do Channel Origem par ao Channel Destino
		destino <- <-origem
	}
}

// Multiplexar : misturar (mensagens) em um único canal
func juntar(entrada1, entrada2 <-chan string) <-chan string {

	canal := make(chan string)
	go encaminhar(entrada1, canal)
	go encaminhar(entrada2, canal)
	return canal
}

func main() {
	canalM := juntar(
		html.Titulo("https://www.cod3r.com.br", "https://www.madeiramadeira.com.br"),
		html.Titulo("https://www.udemy.com", "https://www.uol.com.br"),
	)
	fmt.Println(<-canalM, "|", <-canalM)
	fmt.Println(<-canalM, "|", <-canalM)
}
