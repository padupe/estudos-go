package main

import "fmt"

// func NOME FUNCÃO (PARÂMETRO 1, PARAMÊTRO 2 TIPO PARÂMETRO) (TIPO DE RETORNO 1, TIPO DE RETORNO 2, TIPO DE RETORNO 3)
func compras(trab1, trab2 bool) (bool, bool, bool) {
	comprarTv50 := trab1 && trab2
	comprarTv32 := trab1 != trab2 // OU EXCLUSIVO
	comprarSorvete := trab1 || trab2

	return comprarTv50, comprarTv32, comprarSorvete
}

func main() {
	tv50, tv32, sorvete := compras(true, true)
	fmt.Printf("Tv50: %t, Tv 32: %t, Sorvete: %t, Saudável: %t",
		tv50, tv32, sorvete, !sorvete)
}
