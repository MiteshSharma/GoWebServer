package api

import (
	"github.com/MiteshSharma/project/common/app"
	"github.com/MiteshSharma/project/core/logger"
	"github.com/gorilla/mux"

	"github.com/MiteshSharma/project/core/metrics"
	"github.com/MiteshSharma/project/model"
)

type CommonAPI struct {
	MainRouter *mux.Router
	Config     *model.Config
	Metrics    metrics.Metrics
	Log        logger.Logger
	Router     *Router
	App        *app.App
}

func NewCommonAPI(router *mux.Router, serverParam *model.ServerParam) *CommonAPI {
	api := &CommonAPI{
		MainRouter: router,
		Config:     serverParam.Config,
		Metrics:    serverParam.Metrics,
		Log:        serverParam.Logger,
		Router:     &Router{},
		App:        app.NewApp(serverParam),
	}
	api.setupRoutes()
	return api
}
