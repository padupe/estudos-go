package main

import (
	"fmt"
	"time"
)

func rotina(c chan int) {
	time.Sleep(time.Second)
	c <- 1 // Operação Bloqueante, pois espera-se que algum dado seja enviado para o Channel
	fmt.Println("Só depois que o dado for lido!")
}

func main() {
	c := make(chan int) // Channel sem Buffer [Assunto que ainda será abordado]

	go rotina(c)

	fmt.Println(<-c) // Operação Bloqueante, pois só vou receber o dado, quando ele estiver 'pronto'
	fmt.Println("Dado foi lido com Sucesso!")

	fmt.Println(<-c)                     // Aqui será gerado um Deadlock, pois não existem mais dados a serem recebidos e o código irá falhar
	fmt.Println("Fim do Processamento!") // Nunca chegará nesta linha :(
}
