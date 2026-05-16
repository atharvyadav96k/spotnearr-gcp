package app

import (
	"os"

	db "github.com/atharvyadav96k/spotnearr-gcp/app/database"
)

func Close() {
	db.Close()
}

func (a *App) InitDatabase() error {
	if err := db.Connect(os.Getenv("DATABASE_URL")); err != nil {
		return err
	}
	a.db = db.Get()
	return nil
}
