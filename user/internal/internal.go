package internal

import (
	"errors"

	"github.com/MiteshSharma/project/core/bus"
	"github.com/MiteshSharma/project/core/logger"
	"github.com/MiteshSharma/project/model"
	"github.com/MiteshSharma/project/user/app"
	"github.com/MiteshSharma/project/user/repository"
)

type InternalInput struct {
	Config *model.Config
	Bus    bus.Bus
	Log    logger.Logger
	App    *app.App
}

func NewInternalInput(repository repository.Repository, external repository.External, serverParam *model.ServerParam) *InternalInput {
	internalInput := &InternalInput{
		Config: serverParam.Config,
		Bus:    serverParam.Bus,
		Log:    serverParam.Logger,
		App:    app.NewApp(repository, external, serverParam),
	}

	internalInput.setup()
	return internalInput
}

func (ii *InternalInput) setup() {
	ii.Bus.AddHandler(model.GET_USER, ii.getUser)
}

func (ii *InternalInput) getUser(msg interface{}) (interface{}, error) {
	userID := msg.(int)
	var user *model.User
	var appErr *model.AppError
	if user, appErr = ii.App.GetUser(userID); appErr != nil {
		return nil, errors.New(appErr.ToJson())
	}
	return *user, nil
}
