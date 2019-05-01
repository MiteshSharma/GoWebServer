package repository

import (
	"github.com/MiteshSharma/project/auth/repository/redisRepository"
	"github.com/MiteshSharma/project/auth/repository/sqlRepository"
	"github.com/MiteshSharma/project/core/metrics"
	"github.com/MiteshSharma/project/core/repository/redis"

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
	sqlAuthRepository := repository.PersistentRepository.AuthRepository.(sqlRepository.AuthRepository)
	repository.AuthRepository = redisRepository.NewAuthRepository(repository.RedisRepository, sqlAuthRepository)
	return repository
}

func (s *PersistentCacheRepository) Auth() AuthRepository {
	return s.AuthRepository
}
