package redis

import (
	"context"
	"fmt"
	"smart-route/pkg/config"

	"github.com/redis/go-redis/v9"
)

// NewRedisClientByConfig 根据配置初始化 Redis 客户端
func NewRedisClientByConfig(cfg config.RedisConfig) *redis.Client {
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: cfg.Password,
		DB:       cfg.DB,
		PoolSize: cfg.PoolSize,
	})
}

// RedisPing 测试 Redis 连接
func RedisPing(client *redis.Client) error {
	ctx := context.Background()
	return client.Ping(ctx).Err()
}
