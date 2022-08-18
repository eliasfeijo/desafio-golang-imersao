package model

import "time"

type BankAccount struct {
	ID        uint64    `json:"id"`
	Number    string    `json:"account_number"`
	CreatedAt time.Time `json:"created_at"`
}
