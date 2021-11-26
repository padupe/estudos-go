package main

import "fmt"

// Crio um 'alias' (apelido)
type nota float64

func (n nota) validaEntreConceito(inicio, fim float64) bool {
	return float64(n) >= inicio && float64(n) <= fim
}

func notaParaConceito(n nota) string {
	if n.validaEntreConceito(9.0, 10.0) {
		return "A"
	} else if n.validaEntreConceito(7.0, 8.99) {
		return "B"
	} else if n.validaEntreConceito(5.0, 6.99) {
		return "C"
	} else if n.validaEntreConceito(3.00, 4.99) {
		return "D"
	} else {
		return "E"
	}
}

func main() {
	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(6.99))
	fmt.Println(notaParaConceito(2.55))
}
