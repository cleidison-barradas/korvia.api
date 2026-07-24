package usecases

import (
	"github.com/cleidison-barradas/korvia.api/internal/services/domain/model"
	"github.com/cleidison-barradas/korvia.api/internal/services/domain/repository"
	"github.com/google/uuid"
)

type GetServiceUseCase struct {
	repo repository.ServiceRepository
}

func NewGetServiceUseCase(repo repository.ServiceRepository) *GetServiceUseCase {
	return &GetServiceUseCase{
		repo: repo,
	}
}

func (uc *GetServiceUseCase) Execute(id uuid.UUID) (*model.Services, error) {
	return uc.repo.GetByID(id)
}