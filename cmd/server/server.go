package server

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/MiteshSharma/project/middleware"
	"github.com/MiteshSharma/project/model"
	"github.com/MiteshSharma/project/setting"

	"github.com/MiteshSharma/project/core/bus"
	"github.com/MiteshSharma/project/core/config"
	"github.com/MiteshSharma/project/core/eventdispatcher"

	"github.com/urfave/negroni"

	"github.com/MiteshSharma/project/auth"
	"github.com/MiteshSharma/project/biEvent"
	"github.com/MiteshSharma/project/common"
	"github.com/MiteshSharma/project/core/bi"
	"github.com/MiteshSharma/project/core/logger"
	"github.com/MiteshSharma/project/core/metrics"
	"github.com/MiteshSharma/project/notification"
	"github.com/MiteshSharma/project/user"
	"github.com/gorilla/mux"
)

type Server struct {
	UserServer         *user.UserServer
	CommonServer       *common.CommonServer
	AuthServer         *auth.AuthServer
	BiServer           *biEvent.BiServer
	NotificationServer *notification.NotificationServer
	Router             *mux.Router
	ServerParam        *model.ServerParam
	httpServer         *http.Server
}

func NewServer(settingData *config.Setting) *Server {
	config := setting.GetConfig()
	logger := logger.NewLogger(config.LoggerConfig)
	bus := bus.NewBus(logger)
	metrics := metrics.NewMetrics()
	eventDispatcher := eventdispatcher.NewEventDispatcher(logger, bus, 10, 2)
	biEventHandler := bi.NewBiEventHandler(eventDispatcher, logger)
	router := mux.NewRouter()

	serverParam := model.NewServerParam(logger, metrics, bus, config, eventDispatcher, biEventHandler)

	commonServer := common.NewCommonServer(router, serverParam)
	authServer := auth.NewAuthServer(router, serverParam)
	userServer := user.NewUserServer(router, serverParam)
	notificationServer := notification.NewNotificationServer(router, serverParam)
	biServer := biEvent.NewBiServer(router, serverParam)

	server := &Server{
		UserServer:         userServer,
		CommonServer:       commonServer,
		AuthServer:         authServer,
		BiServer:           biServer,
		NotificationServer: notificationServer,
		Router:             router,
		ServerParam:        serverParam,
	}

	return server
}

func (s *Server) StartServer() {
	n := negroni.New()
	n.UseFunc(middleware.NewLoggerMiddleware(s.ServerParam.Logger).GetMiddlewareHandler())
	if s.ServerParam.Config.ZipkinConfig.IsEnable {
		n.UseFunc(middleware.NewZipkinMiddleware(s.ServerParam.Logger, "project", s.ServerParam.Config.ZipkinConfig).GetMiddlewareHandler())
	}

	n.UseHandler(s.Router)

	listenAddr := (":" + s.ServerParam.Config.ServerConfig.Port)
	s.ServerParam.Logger.Debug("Staring server", logger.String("address", listenAddr))
	s.httpServer = &http.Server{
		Handler:      n,
		Addr:         listenAddr,
		ReadTimeout:  s.ServerParam.Config.ServerConfig.ReadTimeout * time.Second,
		WriteTimeout: s.ServerParam.Config.ServerConfig.WriteTimeout * time.Second,
	}

	go func() {
		err := s.httpServer.ListenAndServe()
		if err != nil {
			s.ServerParam.Logger.Error("Error starting server ", logger.Error(err))
			return
		}
	}()
}

func (s *Server) StopServer() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	s.httpServer.Shutdown(ctx)

	os.Exit(0)
}
