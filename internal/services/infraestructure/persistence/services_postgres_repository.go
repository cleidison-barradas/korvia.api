package persistence

import (
	"context"
	"fmt"

	"github.com/cleidison-barradas/korvia.api/internal/services/domain/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)


type ServicesPostgresRepository struct {
	pool *pgxpool.Pool
	ctx context.Context
}

func NewServicesPostgresRepository(ctx context.Context, pool *pgxpool.Pool) *ServicesPostgresRepository {
	return &ServicesPostgresRepository{
		pool: pool,
		ctx: ctx,
	}
}

func (r *ServicesPostgresRepository) Create(service *model.Services) (*model.Services, error) {
	query := `
		INSERT INTO services (
			name, enabled, description, price, establishment_id, duration
		) VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, name, created_at
	`

	err := r.pool.QueryRow(
		r.ctx, 
		query,
		service.Name,
		service.Enabled,
		service.Description,
		service.Price.Cents,
		service.EstablishmentID,
		service.Duration.Minutes).Scan(&service.Id, &service.Name, &service.CreatedAt)

		if err != nil {
			return nil, err
		}

	return service, nil
}

func (r *ServicesPostgresRepository) GetAll(establishmentID uuid.UUID) ([]model.Services, error) {
	query := `
	SELECT 
		id, 
		name, 
		enabled, 
		description, 
		price, 
		establishment_id, 
		duration,
		created_at	
	FROM services
	WHERE establishment_id = $1
	`

	rows, err := r.pool.Query(r.ctx, query, establishmentID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var services []model.Services

	for rows.Next() {

		service := model.Services{}

		err := rows.Scan(
			&service.Id,
			&service.Name,
			&service.Enabled,
			&service.Description,
			&service.Price.Cents,
			&service.EstablishmentID,
			&service.Duration.Minutes,
			&service.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		services = append(services, service)
	}
	fmt.Println(services)

	return services, nil
}

func (r *ServicesPostgresRepository) GetByID(id uuid.UUID) (*model.Services, error) {
	return nil, nil
}

func (r *ServicesPostgresRepository) Update(id uuid.UUID, service *model.Services) (*model.Services, error) {
	return nil, nil
}

