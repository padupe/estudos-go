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

	// UPDATE
	stmt, _ := db.Prepare("update usuarios set nome = ? where id = ?")

	stmt.Exec("Maria Eduarda", 1)
	stmt.Exec("João Paulo", 2)

	// DELETE
	stmt2, _ := db.Prepare("delete from usuarios where id = ?")

	stmt2.Exec(3)
}
