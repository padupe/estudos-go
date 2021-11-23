package main

import (
	"fmt"
	"time"
)

// FOR : Única estrutura com repetição em Golang

func main() {

	i := 1

	// IMPORTANTE: Golang não possui While, Do While
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)
	}

	fmt.Println("\nMisturando... ")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Print("Par ")
		}
		fmt.Print("Ímpar ")
	}

	for {
		// LAÇO INFINIT: FOR com { }
		fmt.Println("Para Sempre")
		time.Sleep(time.Second)
	}
}
