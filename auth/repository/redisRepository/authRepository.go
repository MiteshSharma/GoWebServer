package redisRepository

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MiteshSharma/project/auth/repository/sqlRepository"
	"github.com/MiteshSharma/project/core/logger"
	"github.com/MiteshSharma/project/core/repository/redis"
	"github.com/MiteshSharma/project/model"
)

type AuthRepository struct {
	*redis.RedisRepository
	sqlRepository.AuthRepository
}

func NewAuthRepository(redisRepository *redis.RedisRepository, authRepository sqlRepository.AuthRepository) AuthRepository {
	authRedisRepository := AuthRepository{
		RedisRepository: redisRepository,
		AuthRepository:  authRepository,
	}

	return authRedisRepository
}

// SaveSession func is used to save user session object in db
func (ar AuthRepository) CreateSession(session *model.UserSession) *model.StorageResult {
	userAuthKey := fmt.Sprintf("user:%d", session.UserID)
	err := ar.Redis.Set(userAuthKey, session.ToJson(), 0).Err()
	if err != nil {
		ar.Log.Error("Error writing redis for user login token", logger.Error(err))
	}
	return ar.AuthRepository.CreateSession(session)
}

func (ar AuthRepository) UpdateSession(session *model.UserSession) *model.StorageResult {
	userAuthKey := fmt.Sprintf("user:%d", session.UserID)
	err := ar.Redis.Set(userAuthKey, session.ToJson(), 0).Err()
	if err != nil {
		ar.Log.Error("Error writing redis for user login token", logger.Error(err))
	}
	return ar.AuthRepository.UpdateSession(session)
}

func (ar AuthRepository) GetSession(userID int) *model.StorageResult {
	_, err := ar.Redis.Ping().Result()
	if err != nil {
		return ar.AuthRepository.GetSession(userID)
	}
	userAuthKey := fmt.Sprintf("user:%d", userID)
	result, err := ar.Redis.Get(userAuthKey).Result()
	if err != nil {
		return ar.AuthRepository.GetSession(userID)
	}
	if result == "" {
		return ar.AuthRepository.GetSession(userID)
	}
	var session model.UserSession
	err = json.Unmarshal([]byte(result), &session)
	if err != nil {
		return model.NewStorageResult(nil, model.NewAppError(err.Error(), http.StatusInternalServerError))
	}
	return model.NewStorageResult(&session, nil)
}

func (ar AuthRepository) DeleteSession(userID int) *model.StorageResult {
	userAuthKey := fmt.Sprintf("user:%d", userID)
	ar.Redis.Del(userAuthKey)
	return ar.AuthRepository.DeleteSession(userID)
}
