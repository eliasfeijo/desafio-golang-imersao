package controller

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/eliasfeijo/desafio-imersao/repository"
)

type BankAccounts interface {
	CreateBankAccount(w http.ResponseWriter, r *http.Request)
}

type bankAccounts struct {
	repository repository.BankAccountsRepository
}

func NewBankAccounts(repository repository.BankAccountsRepository) BankAccounts {
	return &bankAccounts{repository: repository}
}

type CreateBankAccountPayload struct {
	Number string `json:"account_number"`
}

func (r bankAccounts) CreateBankAccount(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Fatalf("Error reading body: %v", err)
	}

	w.Header().Add("Content-Type", "application/json")

	var payload CreateBankAccountPayload
	json.Unmarshal(body, &payload)

	bankAccount, err := r.repository.CreateBankAccount(payload.Number)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode("Error creating BankAccount")
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(bankAccount)
}
