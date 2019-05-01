package user

import (
	"github.com/MiteshSharma/project/model"
	"github.com/MiteshSharma/project/user/api"
	"github.com/MiteshSharma/project/user/internal"
	"github.com/MiteshSharma/project/user/repository"
	"github.com/gorilla/mux"
)

type UserServer struct {
	UserAPI       *api.UserAPI
	InternalInput *internal.InternalInput
}

func NewUserServer(router *mux.Router, serverParam *model.ServerParam) *UserServer {
	storage := repository.NewPersistentCacheRepository(serverParam.Logger, serverParam.Config, serverParam.Metrics)

	external := repository.NewExternalRepository(serverParam.Logger, serverParam.Config,
		serverParam.Metrics, serverParam.Bus)

	userAPI := api.NewUserAPI(router, storage, external, serverParam)

	internalInput := internal.NewInternalInput(storage, external, serverParam)
	server := &UserServer{
		UserAPI:       userAPI,
		InternalInput: internalInput,
	}

	return server
}
