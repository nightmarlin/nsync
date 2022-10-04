package postgres

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

type DB struct {
	log  *zap.Logger
	pool *pgxpool.Pool
}

func New(l *zap.Logger, pool *pgxpool.Pool) DB {
	return DB{
		log:  l,
		pool: pool,
	}
}
