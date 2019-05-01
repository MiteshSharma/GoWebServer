package notification

import (
	"github.com/MiteshSharma/project/model"
	"github.com/MiteshSharma/project/notification/internal"
	"github.com/gorilla/mux"
)

type NotificationServer struct {
	InternalInput *internal.InternalInput
}

func NewNotificationServer(router *mux.Router, serverParam *model.ServerParam) *NotificationServer {
	internalInput := internal.NewInternalInput(serverParam)
	server := &NotificationServer{
		InternalInput: internalInput,
	}

	return server
}
