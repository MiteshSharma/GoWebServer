package common

import (
	"github.com/MiteshSharma/project/common/api"
	"github.com/MiteshSharma/project/model"
	"github.com/gorilla/mux"
)

type CommonServer struct {
	CommonAPI *api.CommonAPI
}

func NewCommonServer(router *mux.Router, serverParam *model.ServerParam) *CommonServer {
	commonAPI := api.NewCommonAPI(router, serverParam)
	server := &CommonServer{
		CommonAPI: commonAPI,
	}

	return server
}
