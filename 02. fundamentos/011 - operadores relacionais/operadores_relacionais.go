package main

import (
	"fmt"
	"time"
)

func main() {
	// Igual
	fmt.Println("IGUAL -> Strings:", "Banana" == "Banana")

	// Diferente
	fmt.Println("DIFERENTE -> !=", 3 != 2)

	// Maior
	fmt.Println("MENOR: <", 3 < 2)

	// Maior
	fmt.Println("MAIOR: >", 3 > 2)

	// Menor ou Igual
	fmt.Println("MENOR OU IGUAL: <=", 3 <= 2)

	// Maior ou Igual
	fmt.Println("MAIOR OU IGUAL: >=", 3 >= 2)

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	fmt.Println("Datas:", d1 == d2)
	// Resultado esperado -> Datas: true

	fmt.Println("Datas:", d1.Equal(d2))
	// Resultado esperado -> Datas: true

	type Pessoa struct {
		Nome string
	}

	p1 := Pessoa{"João"}
	p2 := Pessoa{"João"}
	p3 := Pessoa{"Jão"}

	fmt.Println("Pessoas são iguais?", p1 == p2)
	// Resultado esperado: Pessoas são iguais? true

	fmt.Println("Pessoas são diferentes?", p1 != p3)
	// Resultado esperado: Pessoas são diferentes? true
}
