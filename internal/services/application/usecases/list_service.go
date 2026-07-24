package usecases

import (
	"github.com/cleidison-barradas/korvia.api/internal/services/domain/model"
	"github.com/cleidison-barradas/korvia.api/internal/services/domain/repository"
	"github.com/google/uuid"
)

type ListServiceUseCase struct {
	repo repository.ServiceRepository
}

func NewListServiceUseCase(repo repository.ServiceRepository) *ListServiceUseCase {
	return &ListServiceUseCase{
		repo: repo,
	}
}

func (uc *ListServiceUseCase) Execute(establishmentID uuid.UUID) ([]model.Services, error) {

	return uc.repo.GetAll(establishmentID)
}