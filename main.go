package main

import (
	"log"
	"net/http"
	"os"

	"github.com/eliasfeijo/desafio-imersao/database"
	"github.com/eliasfeijo/desafio-imersao/routes"
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
	routes.SetupRoutesBankAccounts(router)
	routes.SetupRoutesTransfers(router)

	log.Println("API is running")
	http.ListenAndServe(":4000", router)
}
