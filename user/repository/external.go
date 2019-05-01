package repository

import (
	"github.com/MiteshSharma/project/core/bus"
	"github.com/MiteshSharma/project/core/logger"
	"github.com/MiteshSharma/project/core/metrics"
	"github.com/MiteshSharma/project/model"
	"github.com/MiteshSharma/project/user/repository/external"
)

type External interface {
	CreateUserAuth(user model.User) (model.UserAuth, error)
}

type ExternalRepository struct {
	Log          logger.Logger
	Config       *model.Config
	Metrics      metrics.Metrics
	Bus          bus.Bus
	AuthExternal *external.AuthExternal
}

func NewExternalRepository(log logger.Logger, config *model.Config, metrics metrics.Metrics, bus bus.Bus) *ExternalRepository {
	repository := &ExternalRepository{
		Log:     log,
		Config:  config,
		Metrics: metrics,
		Bus:     bus,
	}

	repository.AuthExternal = external.NewAuthExternal(log, config, metrics, bus)

	return repository
}

func (s *ExternalRepository) CreateUserAuth(user model.User) (model.UserAuth, error) {
	return s.AuthExternal.CreateUserAuth(user)
}
