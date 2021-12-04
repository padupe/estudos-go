// NÃO EXISTE EM GO
package main

import "fmt"

func obterResultado(nota float64) string {
	// Alternativa:
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"

	// Se Go admitisse, teríamos algo parecido com:
	// return nota >= 6 ? "Aprovado" : "Reprovado"
}

func main() {
	fmt.Println(obterResultado(6.5))
}
