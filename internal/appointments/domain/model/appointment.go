package model

import (
	"time"

	"github.com/google/uuid"
)

type AppointmentStatus string

const (
	Created AppointmentStatus = "created"
	Confirmed AppointmentStatus = "confirmed"
	Canceled AppointmentStatus = "canceled"
	Finished AppointmentStatus = "finished"
	Deleted AppointmentStatus = "deleted"
	Pending AppointmentStatus = "pending"
)

type Appointment struct {
	Id uuid.UUID `json:"id"`
	UserId uuid.UUID `json:"user_id"`
	ServiceId uuid.UUID `json:"service_id"`
	CustomerId uuid.UUID `json:"customer_id"`
	EstablishmentId uuid.UUID `json:"establishment_id"`
	Status AppointmentStatus `json:"status"`
	StartAt time.Time `json:"start_at"`
	EndAt time.Time `json:"end_at"`
	Observation string `json:"observation omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}