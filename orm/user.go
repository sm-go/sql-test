package main

import (
	"fmt"
	"net/http"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "All Users")
}

func GetSingleUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Single Users")
}

func CreateUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create Users")
}

func UpdateUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update Users")
}

func DeleteUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete Users")
}
