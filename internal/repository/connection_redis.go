package repository

import (
	"context"
	"fmt"
	"rest/internal/models"
	"time"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func ConnectRedis(logger *zap.SugaredLogger, cfg *models.Config) (*redis.Client, error) {
	return NewRedisCacheDB(logger, cfg)
}

func NewRedisCacheDB(logger *zap.SugaredLogger, cfg *models.Config) (*redis.Client, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	redisUri := fmt.Sprintf("%s:%s", cfg.HostRedis, cfg.PortRedis)

	client := redis.NewClient(&redis.Options{
		Addr:        redisUri,
		DB:          0,
		DialTimeout: 100 * time.Millisecond,
		ReadTimeout: 100 * time.Millisecond,
	})

	if _, err := client.Ping(ctx).Result(); err != nil {
		logger.Error("error while ping redis")
		return nil, err
	}

	// defer client.Close()

	return client, nil
}
