package main

import "fmt"

// Poderia ser declarado assim: func trocaValores (p1, p2 int) (segundo, primeiro int) {
func trocaValores(p1, p2 int) (segundo int, primeiro int) {

	// A variável 'segundo' está recebendo o PRIMEIRO PARÂMETRO
	segundo = p2
	// A variável 'primeiro' está recebendo o SEGUNDO PARÂMETRO
	primeiro = p1

	return // Chamado de Retorno Limpo, pois já nomeei os retornos
}

func main() {

	x, y := trocaValores(2, 3)
	fmt.Println(x, y)
	// Retorno Esperado: 3 2

	// Apesar das variáveis receberem o mesmo nome das que existem na Função trocaValores, não tratam-se das MESMAS variáveis
	segundo, primeiro := trocaValores(7, 1)
	fmt.Println(segundo, primeiro)
}
