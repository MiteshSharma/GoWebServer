package sqlRepository

import (
	"net/http"

	"github.com/MiteshSharma/project/core/repository/sql"
	"github.com/MiteshSharma/project/model"
)

type AuthRepository struct {
	*sql.SqlRepository
}

func NewSqlAuthRepository(sqlRepository *sql.SqlRepository) AuthRepository {
	authRepository := AuthRepository{sqlRepository}

	if (!authRepository.DB.HasTable(&model.UserSession{})) {
		// will append "ENGINE=InnoDB" to the SQL statement when creating table `users`
		authRepository.DB.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&model.UserSession{})
		authRepository.DB.Model(&model.UserSession{}).AddIndex("idx_user_id", "user_id")
	}
	return authRepository
}

func (ar AuthRepository) CreateSession(session *model.UserSession) *model.StorageResult {
	if err := ar.DB.Create(&session).Error; err != nil {
		return model.NewStorageResult(nil, model.NewAppError(err.Error(), http.StatusInternalServerError))
	}
	return model.NewStorageResult(session, nil)
}

func (ar AuthRepository) UpdateSession(session *model.UserSession) *model.StorageResult {
	if err := ar.DB.Model(&model.UserSession{}).Where("user_id = ?", session.UserID).Update("token", session.Token).Error; err != nil {
		return model.NewStorageResult(nil, model.NewAppError(err.Error(), http.StatusInternalServerError))
	}
	return model.NewStorageResult(session, nil)
}

func (ar AuthRepository) GetSession(userID int) *model.StorageResult {
	var session model.UserSession
	if err := ar.DB.Where("user_id = ?", userID).First(&session).Error; err != nil {
		return model.NewStorageResult(nil, model.NewAppError(err.Error(), http.StatusInternalServerError))
	}
	return model.NewStorageResult(&session, nil)
}

func (ar AuthRepository) DeleteSession(userID int) *model.StorageResult {
	if err := ar.DB.Where("user_id = ?", userID).Delete(model.UserSession{}).Error; err != nil {
		return model.NewStorageResult(nil, model.NewAppError(err.Error(), http.StatusInternalServerError))
	}
	return model.NewStorageResult(nil, nil)
}
