package main

import (
	"fmt"
	/*
		OPCIONAL
			m "math", eu referencio o "math" como m e posso chamá-lo da seguinte maneira:
				area := PI * m.Pow(raio, 2)
	*/
	"math"
)

func main() {
	/* para definir uma variável:
	1. const ou var
	2. nome da variável
	3. tipo da variável
		IMPORTANTE: sempre que se declarar uma variável, preciso usá-la
	*/
	const PI float64 = 3.1415

	var raio = 3.2 // tipo (float64 - NÚMERO REAL) inferido pelo compilador

	// Forma reduzida de declarar variável:
	// Este formato é o MAIS recomendado
	area := PI * math.Pow(raio, 2)
	fmt.Println("A área da circunferência é", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	var e, f bool = true, false
	fmt.Println(e, f)

	g, h, i := 2, false, "epa!"
	fmt.Println(g, h, i)
}
