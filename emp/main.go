package main

import (
	"log"
	"net/http"
)

func main() {
	//running the server
	log.Println("Server started on: http://localhost:8080")

	//employee mangement
	EmployeeManagement()

	//products mangement
	ProductManagement()

	//listening the port
	http.ListenAndServe(":8080", nil)
}
