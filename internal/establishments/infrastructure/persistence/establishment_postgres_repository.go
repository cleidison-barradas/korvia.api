package persistence

import (
	"context"

	"github.com/cleidison-barradas/korvia.api/internal/establishments/domain/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)


type EstablishmentPostgresRepository struct {
	pool *pgxpool.Pool
	ctx context.Context
}

func NewEstablishmentPostgresRepository(ctx context.Context, pool *pgxpool.Pool) *EstablishmentPostgresRepository {
	return &EstablishmentPostgresRepository{
		pool: pool,
		ctx: ctx,
	}
}

func (r *EstablishmentPostgresRepository) GetByWabaID(waBaId string) (*model.Establishment, error) {
	query := `
		SELECT id, name, enabled, waba_id, phone_number_id, user_id
		FROM establishments
		WHERE waba_id = $1
	`

	var establishment model.Establishment

	err := r.pool.QueryRow(r.ctx, query, waBaId).Scan(
		&establishment.Id,
		&establishment.Name,
		&establishment.Enabled,
		&establishment.WabaId,
		&establishment.PhoneNumberId,
		&establishment.UserId,
	)

	if err != nil {
		return nil, err
	}

	return &establishment, nil
}

func (r *EstablishmentPostgresRepository) GetByID(id uuid.UUID) (*model.Establishment, error) {
	return nil, nil
}

func (r *EstablishmentPostgresRepository) Create(establishment *model.Establishment) (*model.Establishment, error) {
	query := `
		INSERT INTO establishments (
			name, enabled, waba_id, phone_number_id, user_id
		) VALUES ($1, $2, $3, $4, $5)
		RETURNING id, name
	`

	err := r.pool.QueryRow(
		r.ctx, 
		query, 
		establishment.Name, 
		establishment.Enabled,
		establishment.WabaId,
		establishment.PhoneNumberId,
		establishment.UserId,
	).Scan(&establishment.Id, &establishment.Name)

	if err != nil {
		return nil, err
	}

	return establishment, nil
}

func (r *EstablishmentPostgresRepository) Update(id uuid.UUID, establishment *model.Establishment) (*model.Establishment, error) {
	return nil, nil
}