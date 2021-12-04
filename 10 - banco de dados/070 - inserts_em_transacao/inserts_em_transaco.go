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

	// db.Begin() -> inicia uma transaçaõ no Banco de Dados
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into usuarios(id, nome) values(?,?)")

	stmt.Exec(2000, "Beatriz")
	stmt.Exec(2001, "Carlos")
	// O próximo gerarár um ERRO
	_, err = stmt.Exec(1, "Thiago") // Chave Duplicada no Banco de Dados

	// IMPORTANTE: É sempre importante tratar os erros, para evitar que os dados sejam inseridos de maneira incorreta.

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	// Neste ponto o insert será executado
	tx.Commit()
}
