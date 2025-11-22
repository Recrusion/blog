package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Database struct {
	db *sqlx.DB
}

// Создание подключения к базе данных
func ConnectDatabase(dbDriver, dsn string) (*sqlx.DB, error) {
	db, err := sqlx.Connect(dbDriver, dsn)
	if err != nil {
		return nil, fmt.Errorf("error connection to database: %v", err)
	}
	return db, nil
}
