package app

import (
	"github.com/MiteshSharma/project/auth/repository"
	"github.com/MiteshSharma/project/core/bi"
	"github.com/MiteshSharma/project/core/logger"
	"github.com/MiteshSharma/project/core/metrics"
	"github.com/MiteshSharma/project/model"
)

type App struct {
	Repository     repository.Repository
	External       repository.External
	Config         *model.Config
	Metrics        metrics.Metrics
	Log            logger.Logger
	BiEventHandler bi.EventHandler
	UserSession    *model.UserSession
}

func NewApp(repository repository.Repository, external repository.External, serverParam *model.ServerParam) *App {
	app := &App{
		Repository:     repository,
		External:       external,
		Config:         serverParam.Config,
		Metrics:        serverParam.Metrics,
		Log:            serverParam.Logger,
		BiEventHandler: serverParam.BiEventHandler,
	}
	return app
}
