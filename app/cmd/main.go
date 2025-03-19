package main

import (
	"app/app/internal/cache"
	"app/app/internal/config"
	"github.com/redis/go-redis/v9"
	"log"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Ошибка конфигурации: %v", err)
	}

	//TODO : Инициализация логгера
	//logg, err := logger.New(cfg.LoggerConfig)
	//if err != nil {
	//	log.Fatalf("Ошибка логгера: %v", err)
	//}

	redisClient := redis.Client{}
	//redisClient.Set()
	//redis.SearchStdDev
	//redisClient.Del()

	rediska := cache.NewClient(cfg.RedisCfg, redisClient)

}
