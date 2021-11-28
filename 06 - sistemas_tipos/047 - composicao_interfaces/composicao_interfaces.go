package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type luxuoso interface {
	fazerBaliza()
}

type esprotivoLuxuoso interface {
	esportivo
	luxuoso

	// É possível adicionar novos métodos
}

type bmw7 struct{}

func (b bmw7) ligarTurbo() {
	fmt.Println("Turbo...")
}

func (b bmw7) fazerBaliza() {
	fmt.Println("Baliza...")
}

func main() {
	// Variável 'b' implementa a interface esportivoLuxuoso
	var b esprotivoLuxuoso = bmw7{}
	b.ligarTurbo()
	b.fazerBaliza()
}
