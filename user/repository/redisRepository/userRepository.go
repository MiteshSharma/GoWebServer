package redisRepository

import (
	"github.com/MiteshSharma/project/core/repository/redis"
	"github.com/MiteshSharma/project/user/repository/sqlRepository"
)

type UserRepository struct {
	*redis.RedisRepository
	sqlRepository.UserRepository
}

func NewUserRepository(redisRepository *redis.RedisRepository, userRepository sqlRepository.UserRepository) UserRepository {
	userRedisRepository := UserRepository{
		RedisRepository: redisRepository,
		UserRepository:  userRepository,
	}

	return userRedisRepository
}
