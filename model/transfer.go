package model

import "time"

type Transfer struct {
	ID        uint64    `json:"id"`
	FromId    uint64    `json:"from"`
	ToId      uint64    `json:"to"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}
