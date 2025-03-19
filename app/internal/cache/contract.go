package cache

import (
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
	"time"
)

type Cacher interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Get(ctx context.Context, key string) *redis.StringCmd
	Del(ctx context.Context, keys ...string) *redis.IntCmd
}
