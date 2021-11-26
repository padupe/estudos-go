package main

import "fmt"

// RECURSIVIDADE -> A função chama ela mesma

func fatorial(n int) (int, error) {

	switch {
	case n < 0:
		return -1, fmt.Errorf("Número Inválido: %d", n)
	case n == 0:
		return 1, nil
	default:
		// Estou ignorando o ERRO na linha abaixo com o _ (underline)
		fatorialAnterior, _ := fatorial(n - 1)
		return n * fatorialAnterior, nil
	}
}

func main() {

	resultado, _ := fatorial(5)
	fmt.Println(resultado)

	_, err := fatorial(-4)
	if err != nil {
		fmt.Println(err)
	}

	// Existe maneira melhor para aplicar este código
}
