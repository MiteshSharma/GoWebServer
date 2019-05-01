package redis

import (
	"fmt"

	"github.com/go-redis/redis"

	"github.com/MiteshSharma/project/core/config"
	"github.com/MiteshSharma/project/core/logger"
)

type RedisRepository struct {
	Redis  *redis.Client
	Log    logger.Logger
	Config config.CacheConfig
}

func NewRedisRepository(logger logger.Logger, config config.CacheConfig) *RedisRepository {
	redisRepository := &RedisRepository{
		Log:    logger,
		Config: config,
	}
	redisRepository.Redis = redisRepository.getRedis(config)

	return redisRepository
}

func (s *RedisRepository) getRedis(config config.CacheConfig) *redis.Client {
	var client *redis.Client
	if config.Host != "" {
		client = redis.NewClient(&redis.Options{
			Addr:     s.getRedisURL(config),
			Password: config.Password,
			DB:       0,
		})
	}
	return client
}

func (s *RedisRepository) getRedisURL(config config.CacheConfig) string {
	return fmt.Sprintf("%s:%s", config.Host, config.Port)
}
