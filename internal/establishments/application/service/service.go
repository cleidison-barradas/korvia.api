package service

import (
	"github.com/cleidison-barradas/korvia.api/internal/establishments/domain/model"
	"github.com/cleidison-barradas/korvia.api/internal/establishments/domain/repository"
)

type EstablishmentService struct {
	repo repository.EstablishmentRepository
}

func NewEstablishmentService(repo repository.EstablishmentRepository) *EstablishmentService {
	return &EstablishmentService{
		repo: repo,
	}
}

func (s *EstablishmentService) GetByWaBaId(waBaId string) (*model.Establishment, error) {
	return s.repo.GetByWabaID(waBaId)
}