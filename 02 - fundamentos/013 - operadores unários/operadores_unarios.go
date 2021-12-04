package main

import "fmt"

func main() {
	x := 1
	y := 10

	// Apenas a forma Postfix
	x++ // x += 1 ou x = x + 1
	fmt.Println(x)

	y-- // y -= 1 ou y = y -1
	fmt.Println(y)

	// ESSA SINTAXE NÃO É VÁLIDA:
	// fmt.Println(x == y--)
}
