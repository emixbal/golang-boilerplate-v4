package models

import "time"

type Customer struct {
	ID        int       `json:"id"`
	Name      string    `validate:"required" json:"name" bind:"required"`
	Phone     string    `validate:"required" json:"phone" bind:"required"`
	CreatedAt time.Time `validate:"required" json:"created_at" bind:"required"`
}
