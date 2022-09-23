package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:Smith@2022@tcp(127.0.0.1:3306)/db_sql_test")

	if err != nil {
		log.Fatal(err)
	}

	//single row query
	var name string
	err = db.QueryRow("select name from tbl_user where id = ?", 2).Scan(&name)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(name)
	//end single row query

	defer db.Close()
}
