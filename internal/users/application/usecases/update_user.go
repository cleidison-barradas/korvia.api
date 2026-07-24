package usecases

import (
	"github.com/cleidison-barradas/korvia.api/internal/users/domain/model"
	"github.com/cleidison-barradas/korvia.api/internal/users/domain/repository"
	"github.com/google/uuid"
)


type UpdateUserUseCase struct {
	UserRepository repository.UserRepository
}

func NewUpdateUserUseCase(repo repository.UserRepository) *UpdateUserUseCase {
	return &UpdateUserUseCase{
		UserRepository: repo,
	}
}

func (uc *UpdateUserUseCase) Execute(id uuid.UUID, user *model.User) (*model.User, error) {

	u, err := uc.UserRepository.GetByID(id)

	if err != nil {
		return nil, err
	}

	if u == nil {
		return nil, nil
	}

	return uc.UserRepository.Update(id, user)
}