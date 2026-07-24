package usecases

import (
	"github.com/cleidison-barradas/korvia.api/internal/services/application/dto"
	"github.com/cleidison-barradas/korvia.api/internal/services/domain/model"
	"github.com/cleidison-barradas/korvia.api/internal/services/domain/repository"
)

type CreateServiceUseCase struct {
	repo repository.ServiceRepository
}

func NewCreateServiceUseCase(repo repository.ServiceRepository) *CreateServiceUseCase {
	return &CreateServiceUseCase{
		repo: repo,
	}
}

func (Uc *CreateServiceUseCase) Execute(req dto.CreateServiceRequest) (*model.Services, error) {

	s := model.NewService(&model.NewServicesParams{
		Name: req.Name,
		Price: req.Price,
		Description: req.Description,
		Duration: req.Duration,
		Enabled: req.Enabled,
		EstablishmentID: req.EstablishmentID,
	})

	return Uc.repo.Create(s)
}