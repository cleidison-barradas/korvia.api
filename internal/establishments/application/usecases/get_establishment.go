package usecases

import (
	"github.com/cleidison-barradas/korvia.api/internal/establishments/domain/model"
	"github.com/cleidison-barradas/korvia.api/internal/establishments/domain/repository"
	"github.com/google/uuid"
)

type GetEstablishmentUseCase struct {
	repo repository.EstablishmentRepository
}

func NewGetEstablishmentUseCase(repo repository.EstablishmentRepository) *GetEstablishmentUseCase {
	return &GetEstablishmentUseCase{
		repo: repo,
	}
}

func (uc *GetEstablishmentUseCase) GetByID(id uuid.UUID) (*model.Establishment, error) {
	return uc.repo.GetByID(id)
}