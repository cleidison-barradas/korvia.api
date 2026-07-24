package establishments

import (
	"context"

	"github.com/cleidison-barradas/korvia.api/internal/establishments/application/usecases"
	"github.com/cleidison-barradas/korvia.api/internal/establishments/infrastructure/http"
	"github.com/cleidison-barradas/korvia.api/internal/establishments/infrastructure/persistence"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewModule(ctx context.Context, pool *pgxpool.Pool) *http.EstablishmentHandler {
	createUC := usecases.NewCreateEstablishmentUseCase(persistence.NewEstablishmentPostgresRepository(ctx, pool))
	updateUC := usecases.NewUpdateEstablishmentUseCase(persistence.NewEstablishmentPostgresRepository(ctx, pool))

	return http.NewEstablishmentHandler(&http.EstablishmentHandler{
		CreateUC: createUC,
		UpdateUC: updateUC,
	})
}

