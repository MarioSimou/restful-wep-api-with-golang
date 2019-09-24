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
		fmt.Println("executed twice")
		db.CreateTable(&models.User{})
	}

	// routes
	launchServer(db)

}

func launchServer(db *gorm.DB) {
	router := mux.NewRouter().StrictSlash(true)

	// GET /api/users
	router.HandleFunc("/api/users", routes.RenderRoute("getUsers", db)).Methods("GET")
	router.HandleFunc("/api/users", routes.RenderRoute("postUser", db)).Methods("POST")

	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), router))
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
