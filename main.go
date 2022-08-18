package main

import (
	"log"
	"net/http"
	"os"

	"github.com/eliasfeijo/desafio-imersao/controller"
	"github.com/eliasfeijo/desafio-imersao/database"
	"github.com/eliasfeijo/desafio-imersao/repository"
	"github.com/eliasfeijo/desafio-imersao/service"
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

	repository := repository.NewBankAccounts()
	service := service.NewBankAccounts(repository)
	controller := controller.NewBankAccounts(service)

	router := mux.NewRouter()
	router.HandleFunc("/bank-accounts", controller.CreateBankAccount).Methods("POST")
	// router.HandleFunc("/bank-accounts/transfer", controller.CreateTransfer).Methods("POST")

	log.Println("API is running")
	http.ListenAndServe(":4000", router)
}
