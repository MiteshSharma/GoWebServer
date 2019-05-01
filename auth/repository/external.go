package repository

import (
	"github.com/MiteshSharma/project/auth/repository/external"
	"github.com/MiteshSharma/project/core/bus"
	"github.com/MiteshSharma/project/core/logger"
	"github.com/MiteshSharma/project/core/metrics"
	"github.com/MiteshSharma/project/model"
)

type External interface {
	GetUser(userID int) (model.User, error)
}

type ExternalRepository struct {
	Log          logger.Logger
	Config       *model.Config
	Metrics      metrics.Metrics
	Bus          bus.Bus
	UserExternal *external.UserExternal
}

func NewExternalRepository(log logger.Logger, config *model.Config, metrics metrics.Metrics, bus bus.Bus) *ExternalRepository {
	repository := &ExternalRepository{
		Log:     log,
		Config:  config,
		Metrics: metrics,
		Bus:     bus,
	}

	repository.UserExternal = external.NewUserExternal(log, config, metrics, bus)

	return repository
}

func (s *ExternalRepository) GetUser(userID int) (model.User, error) {
	return s.UserExternal.GetUser(userID)
}
