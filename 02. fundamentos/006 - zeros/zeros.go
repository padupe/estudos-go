package main

import "fmt"

func main() {
	var a int
	var b float64
	var c bool
	var d string
	var e *int // Ponteiro

	fmt.Printf("%v %v %v %v %v", a, b, c, d, e)
	// Retorno esperado: 0 0 false  <nil>
	/*
		Entre false e <nil> existe um espaço em branco
		<nil> equivale a null
	*/

	fmt.Printf("\n%v %v %v %q %v", a, b, c, d, e) // O %q permite que retorne uma string vazia ""
	// Retorno esperado: 0 0 false "" <nil>
}
