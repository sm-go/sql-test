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

	fmt.Println("Successfully connect to the MySQL Database")

	defer db.Close()

	//inserting the data simply
	insert, err := db.Query("INSERT INTO table_name (columns) VALUES (values)")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	fmt.Println("Successfully inserted to the database!")

	//single row query
	var name string
	err = db.QueryRow("select name from tbl_user where id = ?", 2).Scan(&name)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(name)
	//end single row query

	//statement that modify data

	stmt, err := db.Prepare("INSERT INTO tbl_user (name, email, address) VALUES(?,?,?)")

	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec("Dolly", "dolly@test.com", "MWD")
	if err != nil {
		log.Fatal(err)
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)

	// _, err := db.Exec("DELETE FROM users")  // OK
	// _, err := db.Query("DELETE FROM users") // BAD

	//connection close
	defer db.Close()
}
