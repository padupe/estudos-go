// Rodamos o comando:
// $ go get -u github.com/go-sql-driver/mysql

package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Criamos a função exec
// Criamos o ponteiro para db
func exec(db *sql.DB, sql string) sql.Result {

	// O resultado da variável será armazenado em 'result'
	result, err := db.Exec(sql)

	if err != nil {
		// Caso dê erro, o panic finaliza o programa
		panic(err)
	}

	return result
}

func main() {
	//                           usuário:senha
	db, err := sql.Open("mysql", "root:123456@/")

	if err != nil {
		panic(err)
	}

	// Essa execução será realizada no final da Função, neste caso, a Conexão com o Banco será encerrada ao final da Função
	defer db.Close()

	exec(db, "create database if not exists cursogolang")
	exec(db, "use cursogolang")
	exec(db, "drop table if exists usuarios")
	exec(db, `create table usuarios (
		id integer auto_increment,
		nome varchar(80),
		PRIMARY KEY (id)
	)`)
}
