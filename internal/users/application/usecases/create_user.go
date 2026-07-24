package usecases

import (
	"github.com/cleidison-barradas/korvia.api/internal/users/application/dto"
	"github.com/cleidison-barradas/korvia.api/internal/users/domain/model"
	"github.com/cleidison-barradas/korvia.api/internal/users/domain/repository"
)

type CreateUserUseCase struct {
	repo repository.UserRepository
}

func NewCreateUserUseCase(repo repository.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{
		repo: repo,
	}
}

func (uc *CreateUserUseCase) Execute(req dto.CreateUserRequest) (*model.User, error) {

	user := model.NewUser(
		req.FirstName,
		req.LastName,
		req.Email,
		req.Phone,
		req.Password,
		req.Document,
		req.DocumentType,
	)

	return uc.repo.Create(user)
}