package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", helloWorld).Methods("GET")
	myRouter.HandleFunc("/users", GetAllUsers).Methods("GET")
	myRouter.HandleFunc("/users/{name}", GetSingleUsers).Methods("GET")
	myRouter.HandleFunc("/users/{name}/{email}/{phone}/{job}/{address}", CreateUsers).Methods("POST")
	myRouter.HandleFunc("/users/{name}/{email}", UpdateUsers).Methods("PUT")
	myRouter.HandleFunc("/users/{name}", DeleteUsers).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	fmt.Println("Go ORM Testing")

	Connect()
	InitialMigration()

	handleRequests()
}
