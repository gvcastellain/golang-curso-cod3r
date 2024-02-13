package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type usuario struct {
	id   int
	nome string
}

func main() {
	db, err := sql.Open("mysql", "root:123456@/cursogo")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmt, _ := db.Query("select id, nome from usuarios where id > ?", 0) //select parametrizado
	defer stmt.Close()

	var user usuario

	for stmt.Next() {
		stmt.Scan(&user.id, &user.nome)
		fmt.Println(user)
	}
}
