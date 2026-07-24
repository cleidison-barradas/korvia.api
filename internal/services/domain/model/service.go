package model

import (
	"time"

	"github.com/google/uuid"
)

type Money struct {
	Cents int64 `json:"cents"`
}

type TimeDuration struct {
	Minutes int `json:"minutes"`
}

type Services struct {
	Id uuid.UUID `json:"id"`
	Name string `json:"name"`
	Enabled bool `json:"enabled"`
	Description string `json:"description"`
	Price Money `json:"price"`
	EstablishmentID uuid.UUID `json:"establishment_id"`
	Duration TimeDuration `json:"duration"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type NewServicesParams struct {
	Name string
	Price int64
	Description string
	Enabled bool
	Duration int
	EstablishmentID uuid.UUID
}

func NewService(p *NewServicesParams) *Services {
	return &Services{
		Name: p.Name,
		Price: Money{ Cents: p.Price },
		Enabled: p.Enabled,
		Description: p.Description,
		Duration: TimeDuration{ Minutes: p.Duration },
		EstablishmentID: p.EstablishmentID,
	}
}