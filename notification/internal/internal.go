package internal

import (
	"github.com/MiteshSharma/project/core/bus"
	"github.com/MiteshSharma/project/core/logger"
	"github.com/MiteshSharma/project/model"
	"github.com/MiteshSharma/project/notification/app"
)

type InternalInput struct {
	Config *model.Config
	Bus    bus.Bus
	Log    logger.Logger
	App    *app.App
}

func NewInternalInput(serverParam *model.ServerParam) *InternalInput {
	internalInput := &InternalInput{
		Config: serverParam.Config,
		Bus:    serverParam.Bus,
		Log:    serverParam.Logger,
		App:    app.NewApp(serverParam),
	}

	internalInput.setup()
	return internalInput
}

func (ii *InternalInput) setup() {
	ii.Bus.AddHandler(model.SEND_NOTIFICATION, ii.sendNotification)
}

func (ii *InternalInput) sendNotification(msg interface{}) (interface{}, error) {
	notificationData := msg.(model.NotificationData)
	go func(notificationData model.NotificationData) {
		ii.App.SendNotification(notificationData)
	}(notificationData)
	return nil, nil
}
