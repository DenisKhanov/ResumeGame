package postgres

import (
	"embed"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"

	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS // Embedding SQL migration files from the "migrations" directory

// MigrationsUp applies all up migrations in the "migrations" directory.
func MigrationsUp(db *pgxpool.Pool) error {
	migrationsDir := "migrations"
	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("goose.SetDialect: %w", err)
	}

	// Run migrations UP to latest version
	if err := goose.Up(stdlib.OpenDBFromPool(db), migrationsDir); err != nil {
		return fmt.Errorf("failed to apply migrations: %v", err)
	}
	return nil
}
