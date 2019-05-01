package auth

import (
	"github.com/MiteshSharma/project/auth/api"
	"github.com/MiteshSharma/project/auth/repository"
	"github.com/MiteshSharma/project/model"
	"github.com/gorilla/mux"
)

type AuthServer struct {
	AuthAPI *api.AuthAPI
}

func NewAuthServer(router *mux.Router, serverParam *model.ServerParam) *AuthServer {

	storageRepository := repository.NewPersistentCacheRepository(serverParam.Logger, serverParam.Config,
		serverParam.Metrics)
	external := repository.NewExternalRepository(serverParam.Logger, serverParam.Config,
		serverParam.Metrics, serverParam.Bus)

	authAPI := api.NewAuthAPI(router, storageRepository, external, serverParam)
	server := &AuthServer{
		AuthAPI: authAPI,
	}

	return server
}
