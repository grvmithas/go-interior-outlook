package dbsetup

import (
	"encoding/json"
	"fmt"
	"net/http"
	"personal/go-interior-outlook/internal/config"

	// for dbsetup
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// DbInit func
func DbInit() (db *gorm.DB) {
	var err error
	res, err := http.Get("http://localhost:8081/readConfig/database")
	if err != nil {
		panic(err.Error())
	}

	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)

	var paramDb config.Database

	err = decoder.Decode(&paramDb)

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf(paramDb.UserName, "paramdb")

	db, err = gorm.Open("mysql", "root:M!thas8733@/interior_outlook?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	//defer db.Close()
	dBase := db.DB()
	//defer dBase.Close()

	err = dBase.Ping()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Connected to database")

	return
}
