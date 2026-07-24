package usecases

import (
	"github.com/cleidison-barradas/korvia.api/internal/establishments/application/dto"
	"github.com/cleidison-barradas/korvia.api/internal/establishments/domain/model"
	"github.com/cleidison-barradas/korvia.api/internal/establishments/domain/repository"
	"github.com/google/uuid"
)

type CreateEstablishmentUseCase struct {
	repo repository.EstablishmentRepository
}

func NewCreateEstablishmentUseCase(repo repository.EstablishmentRepository) *CreateEstablishmentUseCase {
	return &CreateEstablishmentUseCase{
		repo: repo,
	}
}

func (uc *CreateEstablishmentUseCase) Execute(req dto.CreateEstablishmentRequest) (*model.Establishment, error) {
	
	userId, err := uuid.Parse("bca4aed4-f6f8-4676-9444-04e4e23ad63f")

	if err != nil {
		return nil, err
	}

	establishment := model.NewEstablishment(
		req.Name,
		req.WabaId,
		req.PhoneNumberId,
		userId,
	)

	return uc.repo.Create(establishment)
}








