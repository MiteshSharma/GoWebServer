package app

import (
	"github.com/MiteshSharma/project/core/bi"
	"github.com/MiteshSharma/project/core/logger"
	"github.com/MiteshSharma/project/core/metrics"
	"github.com/MiteshSharma/project/core/mq"
	"github.com/MiteshSharma/project/model"
)

type App struct {
	Config         *model.Config
	Metrics        metrics.Metrics
	Log            logger.Logger
	BiEventHandler bi.EventHandler
	MqAgent        mq.MqAgent
}

func NewApp(serverParam *model.ServerParam) *App {
	app := &App{
		Config:         serverParam.Config,
		Metrics:        serverParam.Metrics,
		Log:            serverParam.Logger,
		BiEventHandler: serverParam.BiEventHandler,
		MqAgent:        mq.GetAgent(serverParam.Config.MqConfig),
	}
	return app
}
