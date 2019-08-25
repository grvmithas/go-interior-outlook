package main

import (
	"fmt"
	"net/http"
	"personal/go-interior-outlook/internal/dbsetup"

	"github.com/julienschmidt/httprouter"
)

// Index1 route
func Index1(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func main() {

	dbsetup.DbInit()

	router := httprouter.New()
	router.GET("/", CreateUser)
	router.GET("/createuser", CreateUser)
	http.ListenAndServe(":8080", router)
}
