package main

import "fmt"

func main() {
	// MAP (Mapas) deve ser inicializado, por isso, a linha abaixo não compila
	//var aprovados map[int]string

	aprovados := make(map[int]string)

	aprovados[12345678900] = "Maria"
	aprovados[74185236911] = "João"
	aprovados[36914752844] = "Ana"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {

		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
		// %s = String | %d = INTEIRO
	}

	// RETORNANDO UM VALOR DO MAP
	fmt.Println(aprovados[74185236911])

	// DELETANDO UM VALOR DO MAP
	delete(aprovados, 36914752844)
	fmt.Println(aprovados)
}
