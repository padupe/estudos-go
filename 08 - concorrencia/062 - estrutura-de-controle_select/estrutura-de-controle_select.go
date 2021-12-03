package main

import (
	"fmt"
	"time"

	html "github.com/padupe/titulos-go"
)

// Função vai buscar o título que for encontrado mais rápido de uma URL
func oMaisRapido(url1, url2, url3 string) string {
	channel1 := html.Titulo(url1)
	channel2 := html.Titulo(url2)
	channel3 := html.Titulo(url3)

	// Estrutura de Controle específica para Concorrência
	select {
	case titulo1 := <-channel1:
		return titulo1
	case titulo2 := <-channel2:
		return titulo2
	case titulo3 := <-channel3:
		return titulo3
	case <-time.After(1000 * time.Millisecond):
		return "Todos perderam, pois demorou demais!"
	default:
		return "Sem resposta ainda!"
	}
}

func main() {
	campeao := oMaisRapido(
		"https://www.cod3r.com.br",
		"https://wwww.madeiramadeira.com.br",
		"https://wwww.ovale.com.br",
	)
	fmt.Println(campeao)
}
