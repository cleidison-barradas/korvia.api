package repository

import (
	"github.com/cleidison-barradas/korvia.api/internal/users/domain/model"
	"github.com/google/uuid"
)

type UserRepository interface {
	Create(user *model.User) (*model.User, error)
	GetByID(id uuid.UUID) (*model.User, error)
	Update(id uuid.UUID, user *model.User) (*model.User, error)
}