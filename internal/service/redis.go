package service

import (
	"context"
	"rest/internal/repository"

	"go.uber.org/zap"
)

type CacheService struct {
	repo   *repository.Repository
	logger *zap.SugaredLogger
}

func NewCacheService(repo *repository.Repository, logger *zap.SugaredLogger) *CacheService {
	return &CacheService{repo: repo, logger: logger}
}

func (s *CacheService) AddCounter(ctx context.Context, key string, value int64) error {
	if err := s.repo.Cache.AddCounter(ctx, key, value); err != nil {
		s.logger.Errorf("error while add counter in service: %v", err)
		return err
	}
	return nil
}

func (s *CacheService) SubCounter(ctx context.Context, key string, value int64) error {
	if err := s.repo.Cache.SubCounter(ctx, key, value); err != nil {
		s.logger.Errorf("error while sub counter in service: %v", err)
		return err
	}
	return nil
}

func (s *CacheService) GetCounter(ctx context.Context, key string) (string, error) {
	value, err := s.repo.Cache.GetCounter(ctx, key)
	if err != nil {
		s.logger.Errorf("error while get counter in service: %v", err)
		return "", err
	}
	return value, nil
}
