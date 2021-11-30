package main

import (
	"fmt"
	"time"
)

// Função para verificar se o número é primo
func isPrimo(num int) bool {
	for i := 2; i < num; i++ {
		// Se o número for divísvel por outro (além dele mesmo e por 1), sua divisão retorna 0.
		// Consequentemente este número NÃO é primo
		if num%i == 0 {
			return false
		}
	}
	return true
}

func primos(n int, c chan int) {
	inicio := 2
	for i := 0; i < n; i++ {
		// Este laço, FOR, irá encontrar a quantidade números primos que eu desejo (n)
		for primo := inicio; ; primo++ {
			// Se o retorno da função isPrimo for true, cai no IF
			if isPrimo(primo) {
				c <- primo // Estou enviando dados para o Channel
				// Somo mais um algarismo na variável INICIO, e rodo novamente o laço, FOR
				inicio = primo + 1
				// O código irá aguaradar 100 milisegundos para continuar a execução
				time.Sleep(time.Millisecond * 100)
				// Se encontrar um número Primo, eu realizo o BREAK, e volto para o FOR
				break
			}
		}
	}
	close(c) // Aqui eu fecho o Channel para não receber mais dados
	// O close é importante para evitar que seja gerado um DEADLOCK
}

func main() {

	// estou criando um Channel com Buffer de 30 posições
	// Assim sendo, iremos receber os 30 primeiros números primos que o nosso código compilar
	canal := make(chan int, 30)
	// O método cap retorna a capacidade do Channel
	go primos(cap(canal), canal)
	for primo := range canal {
		fmt.Printf("%d ", primo)
	}
	fmt.Println("Fim!")
}
