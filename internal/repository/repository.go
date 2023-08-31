package repository

import (
	"context"
	"rest/internal/models"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type ICounterCacheRepo interface {
	AddCounter(ctx context.Context, key string, value int64) error
	SubCounter(ctx context.Context, key string, value int64) error
	GetCounter(ctx context.Context, key string) (string, error)
}

type IUserSQL interface {
	CreateUser(user models.User) error
	GetUser(user models.User) (models.User, error)
	UpdateUser(user models.User) error
	DeleteUser(user models.User) error
}

type Repository struct {
	Cache ICounterCacheRepo
	User  IUserSQL
}

func NewRepository(db *sqlx.DB, redis *redis.Client, logger *zap.SugaredLogger, ctx context.Context) *Repository {
	return &Repository{
		Cache: NewRedisCache(redis, logger),
		User:  NewUserSQL(db, logger, ctx),
	}
}
