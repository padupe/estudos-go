package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// Struc de Usuário
type Usuario struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
}

// UsuarioHandler analisa a Request e delega para a função responsável
func UsuarioHandler(res http.ResponseWriter, req *http.Request) {
	// Será realizado um parce, utilizando '/usuarios/' como prefixo
	stringId := strings.TrimPrefix(req.URL.Path, "/usuarios/")

	id, _ := strconv.Atoi(stringId)

	switch {
	case req.Method == "GET" && id > 0:
		usuarioPorID(res, req, id)
	case req.Method == "GET":
		usuarioTodos(res, req)
	default:
		// StatusCode: 404
		res.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(res, "Erro na Requisição!")
	}
}

func usuarioPorID(res http.ResponseWriter, req *http.Request, id int) {
	// Conexão com o Banco de Dados
	dataBase, err := sql.Open("mysql", "root:123456@/cursogolang")
	if err != nil {
		// Evitar de informar o erro detalhadamente para o usuário final
		log.Fatal(err)
	}
	defer dataBase.Close()

	// cad: Cadastro
	var cad Usuario

	// Consulta que busca apenas uma linha -> QueryRow
	dataBase.QueryRow("select id, nome from usuarios where id = ?", id).Scan(&cad.ID, &cad.Nome)

	// Converter o resultado da busca para o Formato JSON
	usuarioJSON, _ := json.Marshal(cad)

	// Informo que o modelo de retorno/resposta é um JSON
	res.Header().Set("Content-Type", "application/json")

	// Resposta da Requisição
	fmt.Fprintf(res, string(usuarioJSON))
}

func usuarioTodos(res http.ResponseWriter, req *http.Request) {
	// Conexão com o Banco de Dados
	dataBase, err := sql.Open("mysql", "root:123456@/cursogolang")
	if err != nil {
		// Evitar de informar o erro detalhadamente para o usuário final
		log.Fatal(err)
	}
	defer dataBase.Close()

	rows, _ := dataBase.Query("select id, nome from usuarios")
	// Fechamos também esta consulta, assim como a conexão com o Banco de Dados
	defer rows.Close()

	// Criamos um Slice do grupo de Usuarios
	var todosUsuarios []Usuario
	// Vamos percorrer o resultado obtido
	for rows.Next() {
		// var usuario to tipo struct Usuario
		var usuario Usuario
		// Mapeamos o ID e Nome
		rows.Scan(&usuario.ID, &usuario.Nome)
		// todosUsuarios recebe um append de usuario (criado na linha 83)
		todosUsuarios = append(todosUsuarios, usuario)
	}

	// Converter o resultado da busca para o Formato JSON
	listaUsuariosJSON, _ := json.Marshal(todosUsuarios)

	// Informo que o modelo de retorno/resposta é um JSON
	res.Header().Set("Content-Type", "application/json")

	// Resposta da Requisição
	fmt.Fprintf(res, string(listaUsuariosJSON))
}
