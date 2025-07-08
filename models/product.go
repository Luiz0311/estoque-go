package models

import "time"

type Product struct {
	ID        uint       `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Amount    int        `json:"amount"`
	Price     float64    `json:"price"`
	Name      string     `json:"name"`
	Type      string     `json:"type"`
	EANCode   string     `json:"ean_code"`
	Available *bool      `json:"available"`
}
