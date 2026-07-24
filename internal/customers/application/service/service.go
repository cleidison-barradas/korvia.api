package service

import (
	"github.com/cleidison-barradas/korvia.api/internal/customers/domain/model"
	"github.com/cleidison-barradas/korvia.api/internal/customers/domain/repository"
)

type CustomerService struct {
	repo repository.CustomerRepository
}

func NewCustomerService(repo repository.CustomerRepository) *CustomerService {
	return &CustomerService{
		repo: repo,
	}
}

func(s *CustomerService) GetByPhone(phone string) (*model.Customer, error) {
	return s.repo.GetByPhone(phone)
}

func (s *CustomerService) CreateCustomer(customer *model.Customer) (*model.Customer, error) {
	return s.repo.Create(customer)
}