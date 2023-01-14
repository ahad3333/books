package storage_testing

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"

	"add/config"
	"add/storage/postgres"
)

var (
	bookRepo *postgres.BookRepo
	categoryRepo *postgres.CategoryRepo
)

func TestMain(m *testing.M) {

	cfg := config.Load()

	config, err := pgxpool.ParseConfig(fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase,
	))

	if err != nil {
		panic(err)
	}

	config.MaxConns = cfg.PostgresMaxConn

	pool, err := pgxpool.ConnectConfig(context.Background(), config)

	if err != nil {
		panic(err)
	}

	bookRepo = postgres.NewBookRepo(pool)

	os.Exit(m.Run())
}
