package usecases

import (
	"github.com/cleidison-barradas/korvia.api/internal/services/domain/model"
	"github.com/cleidison-barradas/korvia.api/internal/services/domain/repository"
	"github.com/google/uuid"
)

type UpdateServiceUseCase struct {
	repo repository.ServiceRepository
}	

func NewUpdateServiceUseCase(repo repository.ServiceRepository) *UpdateServiceUseCase {
	return &UpdateServiceUseCase{
		repo: repo,
	}
}

func (uc *UpdateServiceUseCase) Execute(id uuid.UUID, service *model.Services) (*model.Services, error) {

	return uc.repo.Update(id, service)
}