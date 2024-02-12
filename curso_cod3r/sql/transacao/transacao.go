package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@/cursogo")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	transacao, _ := db.Begin()
	stmt, _ := transacao.Prepare("insert into usuarios(id, nome) values(?,?)")

	stmt.Exec(2000, "Gabriel")
	stmt.Exec(2001, "Isa")
	stmt.Exec(2002, "Mario")
	stmt.Exec(2003, "Doug")

	if err != nil {
		transacao.Rollback()
		log.Fatal(err)
	}

	transacao.Commit()
}
