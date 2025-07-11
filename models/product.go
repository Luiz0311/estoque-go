package models

import "time"

type Product struct {
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	Type       string     `json:"type"`
	EANCode    string     `json:"ean_code"`
	Amount     int        `json:"amount"`
	Price      float64    `json:"price"`
	TotalValue float64    `json:"total_value"`
	Available  bool       `json:"available"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty"`
}
