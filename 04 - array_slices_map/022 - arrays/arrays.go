package main

import "fmt"

func main() {
	// Estruturas Homogêneas (Mesmo tipo de Dado) e Estática (O tamanho não muda/Fixo)

	var notas [3]float64

	fmt.Println(notas)
	// Retorno esperado: [0 0 0]

	notas[0], notas[1], notas[2] = 7.8, 4.3, 9.1

	fmt.Println(notas)
	// Retorno esperado: [ 7.8 4.3 9.1]

	total := 0.0

	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}

	media := total / float64(len(notas))

	fmt.Printf("Média %.2f\n", media) // -> o %.2f indica que é uma variável to tipo FLOAT e que desejo exibir 2 casas decimais
	// Retorno esperado: Média 7.07
}
