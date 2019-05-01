package repository

import (
	"github.com/MiteshSharma/project/core/metrics"
	"github.com/MiteshSharma/project/user/repository/sqlRepository"

	"github.com/MiteshSharma/project/core/logger"
	"github.com/MiteshSharma/project/core/repository/sql"
	"github.com/MiteshSharma/project/model"
)

type PersistentRepository struct {
	SqlRepository *sql.SqlRepository
	Log           logger.Logger
	Config        *model.Config
	Metrics       metrics.Metrics

	UserRepository UserRepository
}

func NewPersistentRepository(log logger.Logger, config *model.Config, metrics metrics.Metrics) *PersistentRepository {
	repository := &PersistentRepository{
		Log:     log,
		Config:  config,
		Metrics: metrics,
	}

	repository.SqlRepository = sql.NewSqlRepository(log, config.DatabaseConfig)
	repository.UserRepository = sqlRepository.NewSqlUserRepository(repository.SqlRepository)

	return repository
}

func (s *PersistentRepository) User() UserRepository {
	return s.UserRepository
}

func (s *PersistentRepository) Close() error {
	return s.SqlRepository.Close()
}
