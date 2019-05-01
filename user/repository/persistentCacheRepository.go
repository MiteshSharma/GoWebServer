package repository

import (
	"github.com/MiteshSharma/project/core/metrics"
	"github.com/MiteshSharma/project/core/repository/redis"
	"github.com/MiteshSharma/project/user/repository/redisRepository"
	"github.com/MiteshSharma/project/user/repository/sqlRepository"

	"github.com/MiteshSharma/project/core/logger"
	"github.com/MiteshSharma/project/model"
)

type PersistentCacheRepository struct {
	*PersistentRepository
	RedisRepository *redis.RedisRepository
}

func NewPersistentCacheRepository(log logger.Logger, config *model.Config, metrics metrics.Metrics) *PersistentCacheRepository {
	repository := &PersistentCacheRepository{
		PersistentRepository: NewPersistentRepository(log, config, metrics),
	}

	repository.RedisRepository = redis.NewRedisRepository(log, config.CacheConfig)
	sqlUserRepository := repository.PersistentRepository.UserRepository.(sqlRepository.UserRepository)
	repository.UserRepository = redisRepository.NewUserRepository(repository.RedisRepository, sqlUserRepository)
	return repository
}

func (s *PersistentCacheRepository) User() UserRepository {
	return s.UserRepository
}
