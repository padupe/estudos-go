package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type usuario struct {
	id   int
	nome string
}

func main() {
	db, err := sql.Open("mysql", "root:123456@/cursogolang")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, _ := db.Query("select id, nome from usuarios where id < ?", 3)
	defer rows.Close()

	for rows.Next() {
		var user usuario
		// & -> referência para o user
		rows.Scan(&user.id, &user.nome)
		fmt.Println(user)
	}
}
