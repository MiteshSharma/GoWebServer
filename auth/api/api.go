package api

import (
	"github.com/MiteshSharma/project/core/logger"
	"github.com/gorilla/mux"

	"github.com/MiteshSharma/project/auth/app"
	"github.com/MiteshSharma/project/auth/repository"
	"github.com/MiteshSharma/project/core/metrics"
	"github.com/MiteshSharma/project/model"
)

type AuthAPI struct {
	MainRouter *mux.Router
	Config     *model.Config
	Metrics    metrics.Metrics
	Log        logger.Logger
	Router     *Router
	App        *app.App
}

func NewAuthAPI(router *mux.Router, repository repository.Repository, external repository.External,
	serverParam *model.ServerParam) *AuthAPI {

	api := &AuthAPI{
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
