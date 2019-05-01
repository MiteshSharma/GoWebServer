package internal

import (
	"github.com/MiteshSharma/project/biEvent/app"
	"github.com/MiteshSharma/project/core/bus"
	"github.com/MiteshSharma/project/core/logger"
	"github.com/MiteshSharma/project/model"
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
	ii.Bus.AddHandler(model.SEND_BI_BATCH_EVENTS, ii.handleBiEventBatch)
}

func (ii *InternalInput) handleBiEventBatch(msg interface{}) (interface{}, error) {
	events := msg.([]map[string]interface{})

	if events != nil {
		ii.Log.Debug("Received events to handle.", logger.Int("size", len(events)))
	}
	ii.App.HandleBIBatch(events)

	return nil, nil
}
