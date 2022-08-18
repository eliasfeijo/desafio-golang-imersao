package controller

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/eliasfeijo/desafio-imersao/service"
)

type Transfers interface {
	CreateTransfer(w http.ResponseWriter, r *http.Request)
}

type transfers struct {
	service service.Transfers
}

func NewTransfers(service service.Transfers) Transfers {
	return &transfers{service: service}
}

type CreateTransferPayload struct {
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
}

func (t transfers) CreateTransfer(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("Error reading body: %v", err)
	}

	var payload CreateTransferPayload
	json.Unmarshal(body, &payload)

	transfer, err := t.service.CreateTransfer(payload.From, payload.To, payload.Amount)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode("Error creating Transfer")
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(transfer)
}
