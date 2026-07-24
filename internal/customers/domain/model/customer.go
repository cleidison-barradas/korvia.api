package model

import (
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	Id uuid.UUID `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Phone string `json:"phone"`
	Email string `json:"email,omitempty"`
	Birthday string `json:"birthday,omitempty"`
	EstablishmentId uuid.UUID `json:"establishment_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func NewCustomer(firstName string, lastName string, phone string, establishmentId uuid.UUID) *Customer {

	return &Customer{
		FirstName: firstName,
		LastName: lastName,
		Phone: phone,
		EstablishmentId: establishmentId,
		CreatedAt: time.Now(),
		UpdatedAt: nil,
	}
}