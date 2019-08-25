package main

import (
	"fmt"
	"net/http"
	"personal/go-interior-outlook/internal/dbsetup"

	"github.com/julienschmidt/httprouter"
)

// Index route
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome Gaurav!\n")
}

// User model
type User struct {
	UserName     string
	UserPassword string
	Phone        string `gorm:"not null"`
	Email        string `gorm:"type:varchar(100);unique_index"`
	ID           int    `gorm:"AUTO_INCREMENT"`
}

// CreateUser to create table
func CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var user = &User{}
	user.UserName = "gaurav"
	user.Phone = "9999620633"
	user.Email = "grvmithas@gmail.com"
	user.UserPassword = "foo"
	user.ID = 1
	db := dbsetup.DbInit()

	fmt.Fprint(w, "phone !\n", user.Phone)
	if !db.HasTable(&User{}) {
		db.CreateTable(&User{})
	}

	db.Create(user)
	fmt.Fprint(w, "Welcome !\n", user.Email)
}
