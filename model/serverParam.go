package model

import (
	"github.com/MiteshSharma/project/core/bi"
	"github.com/MiteshSharma/project/core/bus"
	"github.com/MiteshSharma/project/core/eventdispatcher"
	"github.com/MiteshSharma/project/core/logger"
	"github.com/MiteshSharma/project/core/metrics"
)

type ServerParam struct {
	Logger          logger.Logger
	Metrics         metrics.Metrics
	Bus             bus.Bus
	Config          *Config
	EventDispatcher *eventdispatcher.EventDispatcher
	BiEventHandler  bi.EventHandler
}

func NewServerParam(logger logger.Logger, metrics metrics.Metrics, bus bus.Bus,
	config *Config, eventdispatcher *eventdispatcher.EventDispatcher,
	biEventHandler bi.EventHandler) *ServerParam {
	err := &ServerParam{
		Logger:          logger,
		Metrics:         metrics,
		Bus:             bus,
		Config:          config,
		EventDispatcher: eventdispatcher,
		BiEventHandler:  biEventHandler,
	}
	return err
}
