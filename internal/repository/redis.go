package repository

import (
	"context"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type RedisCache struct {
	client *redis.Client
	logger *zap.SugaredLogger
}

func NewRedisCache(client *redis.Client, logger *zap.SugaredLogger) *RedisCache {
	return &RedisCache{client: client, logger: logger}
}

func (r *RedisCache) AddCounter(ctx context.Context, key string, value int64) error {
	err := r.client.IncrBy(ctx, key, value).Err()
	if err != nil {
		r.logger.Errorf("error while add counter: %v", err)
		return err
	}
	return nil
}

func (r *RedisCache) SubCounter(ctx context.Context, key string, value int64) error {
	if err := r.client.DecrBy(ctx, key, value).Err(); err != nil {
		r.logger.Errorf("error while sub counter: %v", err)
		return err
	}
	return nil
}

func (r *RedisCache) GetCounter(ctx context.Context, key string) (string, error) {
	result, err := r.client.Get(ctx, key).Result()
	if err != nil {
		r.logger.Errorf("error while get counter: %v", err)
		return "", err
	}
	return result, nil
}
