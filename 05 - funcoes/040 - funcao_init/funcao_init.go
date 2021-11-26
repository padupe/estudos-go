package main

import "fmt"

// POR CONVENÇÃO o Go inicializa a compilação pela função Init, mesmo ela não sendo "chamada"
func init() {
	fmt.Println("Inicializando....")
}

func main() {
	fmt.Println("Function Main executada!")
}
