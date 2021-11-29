package main

// Comando rodado no terminal para instalar o pacote: $ go get -u github.com/padupe/calculo-area

// ATENÇÃO: É importante instalar pacotes de fontes confiáveis e que garantam a continuidade dos mesmoa

import (
	"fmt"

	area "github.com/padupe/calculo-area"
)

func main() {
	fmt.Println(area.Circ(6.0))
	// Retorno esperado: 113.0976

	fmt.Println(area.Rect(5.0, 2.0))
	// Retorno esperado: 10

	// A função abaixo não é executada, pois no pacoe ela é PRIVADA
	// fmt.Println(area._TrianguloEq(5.0, 2.0))
}
