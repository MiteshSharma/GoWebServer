package external

import (
	"github.com/MiteshSharma/project/core/bus"
	"github.com/MiteshSharma/project/core/logger"
	"github.com/MiteshSharma/project/core/metrics"
	"github.com/MiteshSharma/project/model"
)

type UserExternal struct {
	Log     logger.Logger
	Config  *model.Config
	Metrics metrics.Metrics
	Bus     bus.Bus
}

func NewUserExternal(log logger.Logger, config *model.Config, metrics metrics.Metrics, bus bus.Bus) *UserExternal {
	repository := &UserExternal{
		Log:     log,
		Config:  config,
		Metrics: metrics,
		Bus:     bus,
	}

	return repository
}

func (ue UserExternal) GetUser(userID int) (model.User, error) {
	userVal, err := ue.Bus.Publish(model.GET_USER, userID)
	return userVal.(model.User), err
}
