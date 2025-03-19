package main

import (
	"app/internal/cache"
	"app/internal/config"
	"context"
	"github.com/redis/go-redis/v9"
	"log"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Ошибка конфигурации: %v", err)
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisCfg.Addr,
		Password: cfg.RedisCfg.Password,
		DB:       cfg.RedisCfg.DB,
	})

	if err = redisClient.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("Не удалось подключиться к Redis: %v", err)
	}

	cache := cache.NewClient(redisClient)

}
