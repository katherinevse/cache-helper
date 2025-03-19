package cache

import (
	"app/app/internal/config"
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

// TODO в контракт интерфейс
type CacheClientInterface interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Get(ctx context.Context, key string) *redis.StringCmd
	Del(ctx context.Context, keys ...string) *redis.IntCmd
}

type RedisCacheClient struct {
	client CacheClientInterface
}

func NewClient(cfg config.RedisConfig, client CacheClientInterface) *RedisCacheClient {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	if err := redisClient.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("Не удалось подключиться к Redis: %v", err)
	}

	return &RedisCacheClient{client: client}
}

func (r *RedisCacheClient) Set(ctx context.Context, key string, value interface{}) error {
	// Вызов Set() и проверка ошибки
	err := r.client.Set(ctx, key, value, 0).Err()
	if err != nil {
		return err
	}

	return nil

}
