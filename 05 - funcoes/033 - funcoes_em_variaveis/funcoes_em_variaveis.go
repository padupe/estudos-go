package main

import "fmt"

var soma = func(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(soma(25, 15))

	subtraco := func(a, b int) int {
		return a - b
	}

	fmt.Println(subtraco(57, 17))
}
