package biEvent

import (
	"github.com/MiteshSharma/project/biEvent/api"
	"github.com/MiteshSharma/project/biEvent/internal"
	"github.com/MiteshSharma/project/model"
	"github.com/gorilla/mux"
)

type BiServer struct {
	BiAPI         *api.BiAPI
	InternalInput *internal.InternalInput
}

func NewBiServer(router *mux.Router, serverParam *model.ServerParam) *BiServer {
	biAPI := api.NewBiAPI(router, serverParam)
	internalInput := internal.NewInternalInput(serverParam)
	server := &BiServer{
		BiAPI:         biAPI,
		InternalInput: internalInput,
	}

	return server
}
