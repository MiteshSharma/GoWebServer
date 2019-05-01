package external

import (
	"github.com/MiteshSharma/project/core/bus"
	"github.com/MiteshSharma/project/core/logger"
	"github.com/MiteshSharma/project/core/metrics"
	"github.com/MiteshSharma/project/model"
)

type AuthExternal struct {
	Log     logger.Logger
	Config  *model.Config
	Metrics metrics.Metrics
	Bus     bus.Bus
}

func NewAuthExternal(log logger.Logger, config *model.Config, metrics metrics.Metrics, bus bus.Bus) *AuthExternal {
	repository := &AuthExternal{
		Log:     log,
		Config:  config,
		Metrics: metrics,
		Bus:     bus,
	}

	return repository
}

func (ae AuthExternal) CreateUserAuth(user model.User) (model.UserAuth, error) {
	userVal, err := ae.Bus.Publish(model.CREATE_AUTH, user)
	return userVal.(model.UserAuth), err
}
