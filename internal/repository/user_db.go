package repository

import (
	"context"
	"database/sql"
	"errors"
	"rest/internal/models"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type UserSQL struct {
	db     *sqlx.DB
	logger *zap.SugaredLogger
	ctx    context.Context
}

func NewUserSQL(db *sqlx.DB, logger *zap.SugaredLogger, ctx context.Context) *UserSQL {
	return &UserSQL{db: db, logger: logger, ctx: ctx}
}

func (r *UserSQL) CreateUser(user models.User) error {
	tx, err := r.db.Begin()
	if err != nil {
		r.logger.Errorf("error while crete: %v", err)
		return err
	}

	query := `
		INSERT INTO users(name, last)
		VALUES ($1, $2)
	`

	if _, err := r.db.ExecContext(r.ctx, query, user.Name, user.Last); err != nil {
		r.logger.Errorf("error while add user to DB: %v", err)
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *UserSQL) GetUser(user models.User) (models.User, error) {
	var result models.User

	query := `
 		SELECT id, name, last 
		FROM users 
		WHERE id=$1
	`

	if err := r.db.QueryRowContext(r.ctx, query, user.Id).Scan(&result.Id, &result.Name, &result.Last); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.User{}, errors.New("user with such ID does not exists")
		}
		r.logger.Errorf("error while get user by user id: %v", err)
		return models.User{}, err
	}

	return result, nil
}

func (r *UserSQL) UpdateUser(user models.User) error {
	tx, err := r.db.Begin()
	if err != nil {
		r.logger.Errorf("error while update: %v", err)
		return err
	}

	query := `
		UPDATE users 
		SET name=$1, last=$2
		WHERE id=$3
	`

	if _, err := r.db.ExecContext(r.ctx, query, user.Name, user.Last, user.Id); err != nil {
		r.logger.Errorf("error while update user by user id: %v", err)
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (r *UserSQL) DeleteUser(user models.User) error {
	tx, err := r.db.Begin()
	if err != nil {
		r.logger.Errorf("error while delete: %v", err)
		return err
	}

	query := `
		DELETE FROM users 
		WHERE id=$1
	`

	if _, err := r.db.ExecContext(r.ctx, query, user.Id); err != nil {
		r.logger.Errorf("error while delete user from DB: %v", err)
		tx.Rollback()
		return err
	}
	return tx.Commit()
}
