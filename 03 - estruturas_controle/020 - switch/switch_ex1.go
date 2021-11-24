package main

import "fmt"

func notaParaConceito(n float64) string {
	var nota = int(n)

	switch nota {
	case 10:
		fallthrough // O código "entra" nesse case e sai executando os demais (abaixo)
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Nota inválida"
	}
}

func main() {
	fmt.Println(notaParaConceito(9.8))
	// Retorno Esperado: A

	fmt.Println(notaParaConceito(6.9))
	// Retorno Esperado: C

	fmt.Println(notaParaConceito(2.1))
	// Retorno Esperado: E

	fmt.Println(notaParaConceito(11))
	// Retorno Esperado: Nota inválida
}
