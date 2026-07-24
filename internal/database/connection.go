package database

import (
	"context"
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
)


func NewPool(ctx context.Context, connString string) (*pgxpool.Pool, error) {

	if connString == "" {
		return nil, fmt.Errorf("connection string is empty")
	}

	pool, err := pgxpool.New(ctx, connString)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "Unable to reach database: %v\n", err)
		return nil, err
	}

	return pool, nil
}

func RunMigrations(connString string) error {
	m, err := migrate.New("file://internal/database/migrations", connString)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create migrations: %v\n", err)
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		fmt.Fprintf(os.Stderr, "Unable to run migrations: %v\n", err)
		return err
	}

	return nil
}