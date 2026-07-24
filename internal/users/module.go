package users

import (
	"context"

	"github.com/cleidison-barradas/korvia.api/internal/users/application/usecases"
	"github.com/cleidison-barradas/korvia.api/internal/users/infraestructure/http"
	"github.com/cleidison-barradas/korvia.api/internal/users/infraestructure/persistence"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewModule(ctx context.Context, pool *pgxpool.Pool) *http.UserHandler {
	createUC := usecases.NewCreateUserUseCase(persistence.NewUserPostgresRepository(ctx, pool))
	updateUC := usecases.NewUpdateUserUseCase(persistence.NewUserPostgresRepository(ctx, pool))
	activateUC := usecases.NewActivateUserUseCase(persistence.NewUserPostgresRepository(ctx, pool))
	
	return http.NewUserHandler(createUC, updateUC, activateUC)
}
	