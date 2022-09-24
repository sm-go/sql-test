package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func ProductManagement() {
	http.HandleFunc("/product", GetAllProduct)
	http.HandleFunc("/product/show", GetSingleProduct)
	http.HandleFunc("/product/create", CreateProduct)
	http.HandleFunc("/product/update", UpdateProduct)
	http.HandleFunc("/product/delete", DeleteProduct)
}

type Product struct {
	Id          int
	Name        string
	Description string
	Category    string
}

func GetAllProduct(w http.ResponseWriter, r *http.Request) {
	db := dbConn()

	selDB, err := db.Query("SELECT * FROM tbl_product ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}
	prod := Product{}
	res := []Product{}
	for selDB.Next() {
		var id int
		var name, description, category string
		err = selDB.Scan(&id, &name, &description, &category)
		if err != nil {
			panic(err.Error())
		}
		prod.Id = id
		prod.Name = name
		prod.Description = description
		prod.Category = category
		res = append(res, prod)
	}
	json.NewEncoder(w).Encode(res)

	defer db.Close()
}

func GetSingleProduct(w http.ResponseWriter, r *http.Request) {
	db := dbConn()

	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM tbl_product WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	prod := Product{}
	for selDB.Next() {
		var id int
		var name, description, category string
		err = selDB.Scan(&id, &name, &description, &category)
		if err != nil {
			panic(err.Error())
		}
		prod.Id = id
		prod.Name = name
		prod.Description = description
		prod.Category = category
	}
	json.NewEncoder(w).Encode(prod)

	defer db.Close()
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	db := dbConn()

	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		category := r.FormValue("category")
		insForm, err := db.Prepare("INSERT INTO tbl_product (name, description, category) VALUES (?, ?, ?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, description, category)
		log.Println("INSERT: Name : " + name + " Description : " + description + " Category : " + category)
	}

	defer db.Close()
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	db := dbConn()

	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		category := r.FormValue("category")
		id := r.FormValue("uid")
		updForm, err := db.Prepare("UPDATE tbl_product SET name=?, description=?, category=? WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		updForm.Exec(name, description, category, id)
		log.Println("UPDATE: PRODUCT WHERE ID = ", id)
	}

	defer db.Close()
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	db := dbConn()

	prod := r.URL.Query().Get("id")
	delQry, err := db.Prepare("DELETE FROM tbl_product WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	delQry.Exec(prod)
	log.Println("DELETED THE PRODUCT WHERE ID = ", prod)

	defer db.Close()
}
