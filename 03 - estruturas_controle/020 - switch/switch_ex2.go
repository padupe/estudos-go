package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch { // Aqui podemos interpretar como: Switch TRUE
	case t.Hour() < 12:
		fmt.Println("Bom Dia!")
	case t.Hour() < 18:
		fmt.Println("Boa Tarde!")
	default:
		fmt.Println("Boa Noite!")
	}
}
