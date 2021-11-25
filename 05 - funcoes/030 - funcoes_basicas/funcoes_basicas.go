package main

import "fmt"

func f1() {
	fmt.Println("Função 1")
}

func f2(p1 string, p2 string) {
	fmt.Printf("Função 2: %s %s\n", p1, p2)
}

func f3() string {
	return "Função 3"
}

func f4(p1, p2 string) string {
	// Sprintf -> Retornar uma Sprint, porém NÃO imprime no Terminal
	return fmt.Sprintf("Função 4: %s %s", p1, p2)
}

func f5() (string, string) {
	return "Retorno 1 / Func5", "Retorno 2 / Func5"
}

func main() {
	f1()
	f2("Paraêmtro 1", "Parâmetro 2")

	resp3, resp4 := f3(), f4("Parâemtro 1/F4", "Parâmetro 2/F4")
	fmt.Println(resp3)
	fmt.Println(resp4)

	respf5A, respf5B := f5()
	fmt.Println("Função 5:", respf5A, respf5B)

	// Poderia ser feito assim, ignorando um parâemtro
	_, resp5D := f5()
	fmt.Println("Outra resposta Função 5:", resp5D)

	resp5C, _ := f5()
	fmt.Println("Mais outra resposta Função 5:", resp5C)
}
