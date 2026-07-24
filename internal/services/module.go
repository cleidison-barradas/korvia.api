package services

import (
	"context"

	"github.com/cleidison-barradas/korvia.api/internal/services/application/usecases"
	"github.com/cleidison-barradas/korvia.api/internal/services/infraestructure/http"
	"github.com/cleidison-barradas/korvia.api/internal/services/infraestructure/persistence"
	"github.com/jackc/pgx/v5/pgxpool"
)


func NewModule(ctx context.Context, pool *pgxpool.Pool) *http.ServiceHandler {
	repo := persistence.NewServicesPostgresRepository(ctx, pool)

	createUC := usecases.NewCreateServiceUseCase(repo)
	listUC := usecases.NewListServiceUseCase(repo)
	getUC := usecases.NewGetServiceUseCase(repo)
	updateUC := usecases.NewUpdateServiceUseCase(repo)


	return http.NewServiceHandler(
		createUC, 
		listUC, 
		getUC, 
		updateUC,
	)
}