package repository

import (
	"github.com/cleidison-barradas/korvia.api/internal/customers/domain/model"
	"github.com/google/uuid"
)

type CustomerRepository interface {
	GetByID(id uuid.UUID) (*model.Customer, error)
	GetByPhone(phone string) (*model.Customer, error)
	Create(customer *model.Customer) (*model.Customer, error)
	Update(id uuid.UUID, customer *model.Customer) (*model.Customer, error)
}