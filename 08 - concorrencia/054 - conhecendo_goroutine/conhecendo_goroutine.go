package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {

	for i := 0; i < qtde; i++ {
		// O metódo abaixo "força" que seja aguardado 1 (um) segundo para cada iteração
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	// fale("Maria", "Por que você não falar comigo?", 3)
	// fale("João", "Só posso falar depois de você Maria!", 1)

	// O comando go cria uma GOROUTINE, e executa o código de maneira independente
	// go fale("Maria", "Ei...", 500)
	// go fale("João", "Opa!", 500)

	// time.Sleep(time.Second * 5)
	// fmt.Println("Fim!")

	go fale("Maria", "Entendi!", 10)
	fale("João", "Parabéns", 5)
}
