package usecases

import (
	"github.com/cleidison-barradas/korvia.api/internal/establishments/domain/model"
	"github.com/cleidison-barradas/korvia.api/internal/establishments/domain/repository"
	"github.com/google/uuid"
)

type UpdateEstablishmentUseCase struct {
	repo repository.EstablishmentRepository
}

func NewUpdateEstablishmentUseCase(repo repository.EstablishmentRepository) *UpdateEstablishmentUseCase {
	return &UpdateEstablishmentUseCase{
		repo: repo,
	}
}

func (uc *UpdateEstablishmentUseCase) Update(id uuid.UUID, establishment *model.Establishment) (*model.Establishment, error) {
	return uc.repo.Update(id, establishment)
}