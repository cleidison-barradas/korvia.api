package repository

import (
	"github.com/cleidison-barradas/korvia.api/internal/establishments/domain/model"
	"github.com/google/uuid"
)

type EstablishmentRepository interface {
	GetByID(id uuid.UUID) (*model.Establishment, error)
	GetByWabaID(wabaID string) (*model.Establishment, error)
	Create(establishment *model.Establishment) (*model.Establishment, error)
	Update(id uuid.UUID, establishment *model.Establishment) (*model.Establishment, error)
}