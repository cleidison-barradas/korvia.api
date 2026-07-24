package repository

import (
	"github.com/cleidison-barradas/korvia.api/internal/services/domain/model"
	"github.com/google/uuid"
)

type ServiceRepository interface {
	GetAll(establishmentID uuid.UUID) ([]model.Services, error)
	GetByID(id uuid.UUID) (*model.Services, error)
	Create(service *model.Services) (*model.Services, error)
	Update(id uuid.UUID, service *model.Services) (*model.Services, error)
}