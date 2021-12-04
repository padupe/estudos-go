package main

import "fmt"

func main() {
	var b byte = 3
	fmt.Println(b)

	i := 3 // inferência de tipo
	i += 3 // Acrescentando valor à variável -> i = i + 3
	i -= 3 // Tirando valor da variável -> i = i - 3
	i /= 2 // Atribuição Divisiva -> i = i / 2
	i *= 2 // Atribuição Multiplicativa -> i = i * 2
	i %= 2 // Atribuição com Módulo -> i = i % 2

	fmt.Println(i)

	// Atribuindo valores a múltiplas variáveis
	x, y := 1, 2
	fmt.Println(x, y)

	// Atribuindo valores de variáveis a outras variáveis
	x, y = y, x
	fmt.Println(x, y)
}
