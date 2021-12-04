package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/usuarios/", UsuarioHandler)
	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

/*
	Para rodar o Código, digite no terminal:
		$ cd 11\ -\ http/075\ -\ integrando_http-sql
		$ go run *.go

	Abra o Browser em:

		- Para retornar todos os usuários
		https://localhost:3000/usuarios/

		- Para retornar um usuário específico
		https://localhost:3000/usuarios/{id}
*/
