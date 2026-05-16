package app

import "github.com/jackc/pgx/v5/pgxpool"

func (a *App) GetDB() *pgxpool.Pool {
	return a.db
}
