package controller

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/eliasfeijo/desafio-imersao/model"
)

type Transfers interface {
	CreateTransfer(w http.ResponseWriter, r *http.Request)
}

func CreateTransfer(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("Error reading body: %v", err)
	}

	var transfer model.Transfer
	json.Unmarshal(body, &transfer)

	//TODO: call service to create Transfer

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(transfer)
}
