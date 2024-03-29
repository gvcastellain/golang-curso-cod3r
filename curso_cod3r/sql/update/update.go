package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@/cursogo")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmt, _ := db.Prepare("update usuarios set nome = ? where id = ?")

	stmt.Exec("NOME ATUALIZADO1", 1)
	stmt.Exec("NOME ATUALIZADO2", 2)

	delete, _ := db.Prepare("delete from usuarios where id = ?")
	delete.Exec(3)
}
