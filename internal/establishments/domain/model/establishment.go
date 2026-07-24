package model

import (
	"time"

	"github.com/google/uuid"
)

type Establishment struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Enabled bool `json:"enabled"`
	WabaId string `json:"waba_id"`
	PhoneNumberId string `json:"phone_number_id"` 
	UserId uuid.UUID `json:"user_id"`
	Description string `json:"description"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func NewEstablishment(name string, wabaId string, phoneNumberId string, userId uuid.UUID) *Establishment {

	return &Establishment{
		Name: name,
		Enabled: false,
		WabaId: wabaId,
		PhoneNumberId: phoneNumberId,
		UserId: userId,
	}
}