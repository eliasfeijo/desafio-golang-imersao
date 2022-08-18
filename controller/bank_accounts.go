package controller

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/eliasfeijo/desafio-imersao/model"
)

type BankAccounts interface {
	CreateBankAccount(w http.ResponseWriter, r *http.Request)
}

func CreateBankAccount(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("Error reading body: %v", err)
	}

	var bankAccount model.BankAccount
	json.Unmarshal(body, &bankAccount)

	//TODO: call service to create BankAccount

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(bankAccount)
}
