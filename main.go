package main

import (
	"log"
	"net/http"

	"github.com/LucasA21/Api-Go-Postgresql/db"
	"github.com/LucasA21/Api-Go-Postgresql/routes"
	"github.com/gorilla/mux"
)

func main() {

	db.DBConeection()

	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)

	log.Println("Server runing in http://localhost:8080/")
	http.ListenAndServe(":8080", r)

}
