package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"models"
	"net/http"
	"os"
	"routes"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	if e := godotenv.Load(); e != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	var users models.Users

	f, e := os.Open(os.Getenv("MOCK_PATH"))
	check(e)
	defer f.Close()

	b, e := ioutil.ReadAll(f)
	check(e)

	if e := json.Unmarshal(b, &users); e != nil {
		log.Fatal(e)
	}

	// routes
	launchServer(&users)

}

func launchServer(users *models.Users) {
	router := mux.NewRouter().StrictSlash(true)

	// GET /api/users
	router.HandleFunc("/api/users", routes.GetUsers)

	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), router))
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
