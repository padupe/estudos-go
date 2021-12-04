package main

import "fmt"

func main() {
	// Print -> imprime na mesma linha
	fmt.Print("Mesma")
	fmt.Print(" linha")
	// Retorno: Mesma linha.>

	// Println -> imprime na próxima linha
	fmt.Println(".>")
	// Retorno: Mesma linha.>
	fmt.Println("Nova Linha")
	// Retorno: Nova Linha

	x := 3.141516

	// fmt.Println("O valor de x é " + x) -> Desta maneira não funciona

	// Sprint retorna da função uma string
	xs := fmt.Sprint(x)
	fmt.Println("O valor de x é " + xs)
	// Retorno: O valor de x é 3.141516

	// Posso retornar o valor apenas mudando a sintaxe, conforme na Aula 003:
	fmt.Println("O valor de x é", x)
	// Retorno: O valor de x é 3.141516

	// Printf -> imprime de maneira formatada, conforme determinado no código
	// O .2f informa que quero apenas duas casas decimais
	fmt.Printf("O valor de x é %.2f", x)
	// Retorno: O valor de x é 3.14

	/*
		%d - Variável do tipo INTEIRO
		%f - Variável do tipo FLOAT
		%s - Variável do tipo STRING
		%t - Variável do tipo BOOLEAN
		%v - Variável para QUASE TODOS os tipos
	*/

	a := 1
	b := 1.9999
	c := false
	d := "opa"

	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	// Retorno: 1 1.999900 2.0 false opa

	fmt.Printf("\n%v %v %v %v", a, b, c, d)
	// Retorno: 1 1.9999 false opa
}
