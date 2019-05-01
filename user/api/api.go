package api

import (
	"github.com/MiteshSharma/project/core/logger"
	"github.com/gorilla/mux"

	"github.com/MiteshSharma/project/core/metrics"
	"github.com/MiteshSharma/project/model"
	"github.com/MiteshSharma/project/user/app"
	"github.com/MiteshSharma/project/user/repository"
)

type UserAPI struct {
	MainRouter *mux.Router
	Config     *model.Config
	Metrics    metrics.Metrics
	Log        logger.Logger
	Router     *Router
	App        *app.App
}

func NewUserAPI(router *mux.Router, repository repository.Repository, external repository.External, serverParam *model.ServerParam) *UserAPI {
	api := &UserAPI{
		MainRouter: router,
		Config:     serverParam.Config,
		Metrics:    serverParam.Metrics,
		Log:        serverParam.Logger,
		Router:     &Router{},
		App:        app.NewApp(repository, external, serverParam),
	}
	api.setupRoutes()
	return api
}
