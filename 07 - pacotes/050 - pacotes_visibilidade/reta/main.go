package main

import "fmt"

func main() {
	p1 := Ponto{2.0, 2.0}
	p2 := Ponto{2.0, 4.0}

	fmt.Print("Catetos: ")
	fmt.Println(catetos(p1, p2))
	fmt.Print("Distância: ")
	fmt.Println(Distancia(p1, p2))

	/*
		Retorno Esperado:	Catetos: 0 2
							Distância: 2
	*/
}

/*
	PARA EXECUTAR ESTE CÓDIGO, NO TERMINAL REALIZE:

	cd 07\ -\ pacotes
	cd 050\ -\ pacotes_visibilidade
	cd reta

	$ go run *.go
*/
