// db/db.go
package db

import (
	"context"
	"fmt"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	pool *pgxpool.Pool
	once sync.Once
)

func Connect(databaseURL string) error {
	var initErr error
	once.Do(func() {
		config, err := pgxpool.ParseConfig(databaseURL)
		if err != nil {
			initErr = fmt.Errorf("failed to parse config: %w", err)
			return
		}

		pool, err = pgxpool.NewWithConfig(context.Background(), config)
		if err != nil {
			initErr = fmt.Errorf("failed to create pool: %w", err)
			return
		}

		if err := pool.Ping(context.Background()); err != nil {
			initErr = fmt.Errorf("failed to ping database: %w", err)
			return
		}
	})
	return initErr
}

func Get() *pgxpool.Pool {
	if pool == nil {
		panic("db: not connected — call db.Connect() first")
	}
	return pool
}

func Close() {
	if pool != nil {
		pool.Close()
	}
}
