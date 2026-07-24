package dto

import (
	"github.com/google/uuid"
)

type CreateServiceRequest struct {
	Name string `json:"name" binding:"required"`
	Enabled bool `json:"enabled"`
	Description string `json:"description"`
	Price int64 `json:"price" binding:"required"`
	EstablishmentID uuid.UUID `json:"establishment_id" binding:"required"`
	Duration int `json:"duration" binding:"required"`
}