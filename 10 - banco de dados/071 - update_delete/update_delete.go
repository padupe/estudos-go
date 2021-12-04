package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@/cursogolang")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, _ := db.Prepare("update usuarios set nome = ? where id = ?")

	stmt.Exec("Maria Eduarda", 1)
	stmt.Exec("João Paulo", 2)

	stmt2, _ := db.Prepare("delete from usuarios where id = ?")

	stmt2.Exec(3)
}

// git commit -m $'feat(10 - banco de dados/071 - update_delete): Adicionado ao Projeto.\nAula 92 - Curso Go (Golang): Explorando a linguagem do Google.'
