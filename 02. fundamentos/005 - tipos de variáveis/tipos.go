package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// NUMÉRICOS INTEIROS
	fmt.Println(1, 2, 1000)
	// Retorno esperado: 1 2 1000

	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))
	// Retorno esperado: Literal inteiro é int

	// SEM SINAL (Só positivos)... uint8 uint16 uint32 uint64
	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))
	// Retorno esperado: O byte é uint8

	// COM SINAL... int8 int16 int32 int65
	i1 := math.MaxInt64

	fmt.Println("O valor máximo do int é", i1)
	// Retorno esperadi: O valor máximo do int é 9223372036854775807

	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))
	// Retorno esperado: O tipo de i1 é int

	var i2 rune = 'a' // Representa um mapeamento da tabela Unicode (int32)
	fmt.Println("O rune é", reflect.TypeOf(i2))
	// Retorno esperado: O rune é int32

	fmt.Println(i2)
	// Retorno esperado: 97

	// NÚMEROS REAIS
	var x float32 = 49.99 // poderia ser escrito: var x = float32(49.99)
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	// Retorno esperado: O tipo de x é float32

	fmt.Println("O tipo do literal 49.99 é", reflect.TypeOf(49.99)) // float64
	// Retorno esperado: O tipo do literal 49.99 é float64

	// BOOLEAN
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	// Retorno esperado: O tipo de bo é bool

	fmt.Println(!bo)
	// Retorno esperado: false

	// STRING (Delimitado por aspas duplas)
	s1 := "Olá, meu nome é Paulo"
	fmt.Println(s1 + "!")
	// Retorno esperado: Olá, meu nome é Paulo!

	fmt.Println("O tamanho da string é", len(s1))
	// Retorno esperado: O tamanho da string é 23

	// STRINGS COM MÚLTIPLAS LINHAS
	s2 := `Olá
	meu
	nome
	é
	Paulo`
	fmt.Println("O tamanho da string é", len(s2)) // Resultado diferente pois conta-se as quebras de linha
	// Retorno esperado: O tamanho da string é 26

	// CHAR NÃO EXISTE
	// var x char = 'b'
	char := 'a'
	fmt.Println("O tipo de char é", reflect.TypeOf(char))
	// Retorno esperado: O tipo de char é int32

	fmt.Println(char)
	// Retorno esperado: 97
}
