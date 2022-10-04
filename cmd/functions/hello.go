package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"

	"github.com/nightmarlin/nsync/core/storage/postgres"
)

func main() {
	ctx := context.Background()
	l, err := zap.NewDevelopment()
	if err != nil {
		panic(fmt.Errorf("failed to init logger: %w", err))
	}

	pool, err := pgxpool.Connect(ctx, os.Getenv("DSN"))
	if err != nil {
		l.Fatal("failed to init db connection pool", zap.Error(err))
	}
	defer pool.Close()

	pg := postgres.New(l, pool)
}
