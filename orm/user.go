package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "root:Smith@2022@tcp(127.0.0.1:3306)/db_sql_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	DB = db

	fmt.Println("Successfully connect to the MySQL Database")

}

type User struct {
	gorm.Model
	Name    string
	Email   string
	Phone   string
	Job     string
	Address string
}

func InitialMigration() {
	// Connect()
	DB.AutoMigrate(&User{}) //auto migrate struct User = users (table)
	fmt.Println("Successfully InitialMigration to the database!")
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "All Users")
	var users []User
	DB.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func GetSingleUsers(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	name := vars["name"]
	var user User
	DB.Where("name = ?", name).Find(&user)
	json.NewEncoder(w).Encode(user)

	fmt.Fprintf(w, "Single Users")
}

func CreateUsers(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]
	phone := vars["phone"]
	job := vars["job"]
	address := vars["address"]
	// fmt.Println(name, email, "DDD", mux.Vars(r))

	//URL : localhost:8080/users/smith/smith@test.com/098893434/Developer/Yangon

	DB.Create(&User{Name: name, Email: email, Phone: phone, Job: job, Address: address})
	fmt.Fprintf(w, "New User Successfully Created.")

}

func UpdateUsers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]
	// phone := vars["phone"]
	// job := vars["job"]
	// address := vars["address"]

	var user User
	DB.Where("name = ?", name).Find(&user)
	user.Email = email

	DB.Save(&user)

	fmt.Fprintf(w, "Users successfully Updated.")
}

func DeleteUsers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	var user User
	DB.Where("name = ?", name).Find(&user)
	DB.Delete(user)
	fmt.Fprintf(w, "Users successfully Deleted.")
}
