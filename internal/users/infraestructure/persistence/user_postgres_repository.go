package persistence

import (
	"context"

	"github.com/cleidison-barradas/korvia.api/internal/users/domain/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserPostgresRepository struct {
	pool *pgxpool.Pool
	ctx context.Context
}

func NewUserPostgresRepository(ctx context.Context, pool *pgxpool.Pool) *UserPostgresRepository {
	return &UserPostgresRepository{
		pool: pool,
		ctx: ctx,
	}
}

func (r *UserPostgresRepository) GetByID(id uuid.UUID) (*model.User, error) {

	query := `
	SELECT 
		id, first_name, last_name, email, phone, enabled,
		role, password_hash, document, document_type
		FROM users
		WHERE id = $1
	`
	user := &model.User{}

	err := r.pool.QueryRow(r.ctx, query, id).Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Phone,
		&user.Enabled,
		&user.Role,
		&user.PasswordHash,
		&user.Document,
		&user.DocumentType,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserPostgresRepository) Create(user *model.User) (*model.User, error) {
	query := `
		INSERT INTO users (
			first_name, last_name, email, phone, enabled,
			role, password_hash, document, document_type
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id, first_name, last_name, email, phone, enabled,
			role, password_hash, document, document_type
	`

	err := r.pool.QueryRow(
		r.ctx, 
		query, 
		user.FirstName, 
		user.LastName,
		user.Email,
		user.Phone,
		user.Enabled,
		user.Role,
		user.PasswordHash,
		user.Document,
		user.DocumentType,
	).Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Phone, &user.Enabled, &user.Role, &user.PasswordHash, &user.Document, &user.DocumentType)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserPostgresRepository) Update(id uuid.UUID,	user *model.User) (*model.User, error) {

	query := `
		UPDATE users
		SET first_name = $2, last_name = $3, document = $4, document_type = $5
		WHERE id = $1
	`

	_, err := r.pool.Exec(r.ctx, query, id, user.FirstName, user.LastName, user.Document, user.DocumentType)

	if err != nil {
		return nil, err
	}

	return user, nil
}