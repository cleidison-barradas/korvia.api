package usecases

import (
	"errors"
	"time"

	"github.com/cleidison-barradas/korvia.api/internal/users/domain/repository"
	"github.com/google/uuid"
)

type ActivateUserUseCase struct {
	repo repository.UserRepository
}

func NewActivateUserUseCase(repo repository.UserRepository) *ActivateUserUseCase {
	return &ActivateUserUseCase{
		repo: repo,
	}
}

func (uc *ActivateUserUseCase) Execute(id uuid.UUID) error {

	user, err := uc.repo.GetByID(id)

	if err != nil {
		return err
	}

	if user == nil {
		return errors.New("user not found")
	}

	if user.Enabled {
		return errors.New("user already enabled")
	}

	user.Enabled = true
	user.VerifiedAt = &time.Time{}

	_, e := uc.repo.Update(id, user)

	if e != nil {
		return e
	}

	return nil
}