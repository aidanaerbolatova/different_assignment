package service

import (
	"context"
	"rest/internal/models"
	"rest/internal/repository"

	"go.uber.org/zap"
)

type ICounterCacheService interface {
	AddCounter(ctx context.Context, key string, value int64) error
	SubCounter(ctx context.Context, key string, value int64) error
	GetCounter(ctx context.Context, key string) (string, error)
}

type IUserService interface {
	CreateUser(user models.User) error
	GetUser(user models.User) (models.User, error)
	UpdateUser(user models.User) error
	DeleteUser(user models.User) error
}

type Service struct {
	Cache ICounterCacheService
	User  IUserService
}

func NewService(repo *repository.Repository, logger *zap.SugaredLogger) *Service {
	return &Service{
		Cache: NewCacheService(repo, logger),
		User:  NewUserService(repo, logger),
	}
}
