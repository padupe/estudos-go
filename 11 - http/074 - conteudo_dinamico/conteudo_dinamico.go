package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func horaCerta(w http.ResponseWriter, r *http.Request) {

	// FORMATANDO DATAS E HORAS EM GOLANG:
	// https://www.alura.com.br/artigos/golang-trabalhando-com-datas?gclid=CjwKCAiAwKyNBhBfEiwA_mrUMvcToTKIDL2OaxixlU1L79TmDkWOycHcbPj172Dt4j2e5PT2ljkiVhoCytoQAvD_BwE
	hora := time.Now().Format("02/01/2006 03:04:05")
	// FPrintf -> escrever no Writer
	fmt.Fprintf(w, "<h1>Hora certa: %s<h1>", hora)
}

func main() {
	http.HandleFunc("/horacerta", horaCerta)
	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

/*
	PARA RODAR, EM SEU TERMINAL:

		$ cd 11\ -\ http/074\ -\ conteudo_dinamico/
		$ go run conteudo_dinamico.go

		No navegador, acesse:
		https://localhost:3000/horacerta
*/
