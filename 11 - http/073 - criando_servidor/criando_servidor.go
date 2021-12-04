package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

/*
	PARAR RODAR O SERVIDOR:
	Pelo terminal, rode os comandos:

		$ cd 11\ -\ http/073\ -\ criando_servidor/
		$ go run criando_servidor_go

	No Browser, acesse:

		http://localhost:3000/

*/
