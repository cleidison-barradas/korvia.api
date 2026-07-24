package persistence

import (
	"context"

	"github.com/cleidison-barradas/korvia.api/internal/customers/domain/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CustomerPostgresRepository struct {
	pool *pgxpool.Pool
	ctx context.Context
}

func NewCustomerPostgresRepository(ctx context.Context, pool *pgxpool.Pool) *CustomerPostgresRepository {
	return &CustomerPostgresRepository{
		pool: pool,
		ctx: ctx,
	}
}

func (r *CustomerPostgresRepository) GetByPhone(phone string) (*model.Customer, error) {
	query := `
	SELECT 
		id, first_name, last_name, phone, email, birthday, establishment_id
		FROM customers
		WHERE phone = $1
	`
	customer := &model.Customer{}

	err := r.pool.QueryRow(r.ctx, query, phone).Scan(customer.Id, customer.FirstName, customer.LastName, customer.Phone)

	if err != nil {
		return nil, nil
	}

	return customer, nil
}

func (r *CustomerPostgresRepository) Create(customer *model.Customer) (*model.Customer, error) {

	query := `
	INSERT INTO customers (
		id, first_name, last_name, phone, email, birthday, establishment_id
	) VALUES ($1, $2, $3, $4, $5, $6, $7)
	 RETURNING id, first_name, last_name, phone
	 `

	c := &model.Customer{}

	err := r.pool.QueryRow(r.ctx, query, customer.Id, customer.FirstName, customer.LastName, customer.Phone, customer.Email, customer.Birthday, customer.EstablishmentId).Scan(&c.Id, &c.FirstName, &c.LastName, &c.Phone)

	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *CustomerPostgresRepository) Update(id uuid.UUID, customer *model.Customer) (*model.Customer, error) {
	return nil, nil
}

func (c *CustomerPostgresRepository) GetByID(id uuid.UUID) (*model.Customer, error) {
	return nil, nil
}