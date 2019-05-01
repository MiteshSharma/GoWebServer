package repository

import "github.com/MiteshSharma/project/model"

type Repository interface {
	Close() error
	Auth() AuthRepository
}

type AuthRepository interface {
	CreateSession(session *model.UserSession) *model.StorageResult
	UpdateSession(session *model.UserSession) *model.StorageResult
	GetSession(userID int) *model.StorageResult
	DeleteSession(userID int) *model.StorageResult
}
