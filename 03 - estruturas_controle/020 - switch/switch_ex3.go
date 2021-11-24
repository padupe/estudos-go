package main

import (
	"fmt"
	"time"
)

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "Inteiro"
	case float32, float64:
		return "Real"
	case string:
		return "String"
	case func():
		return "Função"
	default:
		return "Tipo não identificado!"
	}
}

func main() {
	fmt.Println(tipo(2.3))
	// Retorno Esperado: Real

	fmt.Println(tipo(3))
	// Retorno Esperado: Inteiro

	fmt.Println(tipo("Estudando Golang"))
	// Retorno Esperado: String

	fmt.Println(tipo(func() {}))
	// Retorno Esperado: Função

	fmt.Println(tipo(time.Now()))
	// Retorno Esperado: Tipo não identificado!

}
