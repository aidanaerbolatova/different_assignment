package repository

import (
	"fmt"
	"rest/internal/models"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func ConnectSQL(logger *zap.SugaredLogger, config *models.Config) (*sqlx.DB, error) {
	return NewPostgreSQL(logger, config)
}

func NewPostgreSQL(logger *zap.SugaredLogger, config *models.Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", config.HostSQL, config.PortSQL, config.UsernameSQL, config.DBName, config.PasswordSQL, config.SSLmode))
	if err != nil {
		logger.Errorf("error while connect to DB: %v", err)
		return nil, err
	}

	if err := CreateTable(db); err != nil {
		logger.Errorf("error while create table: %v", err)
		return nil, err
	}

	return db, nil
}

func CreateTable(db *sqlx.DB) error {
	query := `CREATE TABLE IF NOT EXISTS "users"(id SERIAL, name VARCHAR, last VARCHAR);`

	if _, err := db.Exec(query); err != nil {
		return err
	}
	return nil
}
