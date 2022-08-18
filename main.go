package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/eliasfeijo/desafio-imersao/database"
	"github.com/gorilla/mux"
)

func main() {
	db, err := database.ConnectDatabase()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		os.Exit(1)
	}
	defer db.Close()

	err = database.Migrate()
	if err != nil {
		log.Fatalf("Error migrating database: %v", err)
		os.Exit(1)
	}

	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Hello World")
	})

	log.Println("API is running")
	http.ListenAndServe(":4000", router)
}
