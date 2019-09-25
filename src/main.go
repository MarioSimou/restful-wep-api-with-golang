package main

import (
	"fmt"
	"log"
	"models"
	"net/http"
	"os"
	"routes"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

func init() {
	if e := godotenv.Load(); e != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	db, e := gorm.Open("postgres", os.Getenv("POSTGRES_URI")+"?sslmode=disable")
	check(e)

	defer db.Close()

	// creates the table if does not exists
	if !db.HasTable(models.User{}) {
		fmt.Println("user table is created...")
		db.CreateTable(&models.User{})
	}

	// routes
	launchServer(db)

}

func launchServer(db *gorm.DB) {
	router := mux.NewRouter().StrictSlash(true)
	api := router.PathPrefix("/api").Subrouter()

	// GET /api/users
	api.HandleFunc("/users", routes.FetchRoute(routes.GET_USERS, db)).Methods("GET")
	api.HandleFunc("/users/{id:[0-9]+}", routes.FetchRoute(routes.GET_USER, db)).Methods("GET")
	api.HandleFunc("/users", routes.FetchRoute(routes.POST_USER, db)).Methods("POST")
	api.HandleFunc("/users/{id:[0-9]+}", routes.FetchRoute(routes.DELETE_USER,db)).Methods("DELETE")
	api.HandleFunc("/users/{id:[0-9]+}", routes.FetchRoute(routes.PUT_USER, db)).Methods("PUT")

	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), router))
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
