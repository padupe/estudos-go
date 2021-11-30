package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// Padrão Generator
// Google I/O 2012 - Go Concurrency Patterns

// <-chan : Canal somente Leitura

func titulo(urls ...string) <-chan string {
	canal := make(chan string)

	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			regex, _ := regexp.Compile("<title>(.*?)<\\/title>")
			canal <- regex.FindStringSubmatch(string(html))[1]
		}(url)
	}

	return canal
}

func main() {
	titulo1 := titulo("https://www.madeiramadeira.com.br", "https://www.udemy.com")
	titulo2 := titulo("https://www.cod3r.com.br", "https://www.sjc.sp.gov.br")

	fmt.Println("Primeiros colocados:", <-titulo1, "|", <-titulo2)
	fmt.Println("Segundos colocados:", <-titulo1, "|", <-titulo2)
}

// REPO PROF CURSO: https://github.com/cod3rcursos/curso-go/blob/master/concorrencia/generator/generator.go
