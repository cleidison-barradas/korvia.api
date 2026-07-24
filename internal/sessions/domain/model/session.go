package model

import (
	"time"

	"github.com/cleidison-barradas/korvia.api/internal/sessions/application/dto"
)

type Session struct {
	PhoneNumber string `json:"phone_number"`
	State dto.State `json:"state"`
	Context dto.Context `json:"context"`
	History []dto.State `json:"history"`
	LastOptions map[string]dto.Option `json:"last_options"`
	LastText string `json:"last_text"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type NewSessionParams struct {
	PhoneNumber string 
	Context dto.Context
}

func NewSession(p *NewSessionParams) *Session {
	return &Session{
		Context: p.Context,
		State: dto.STATE_MENU,
		PhoneNumber: p.PhoneNumber,
		CreatedAt: time.Now(),
	}
}

