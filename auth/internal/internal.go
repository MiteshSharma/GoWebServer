package internal

import (
	"errors"

	"github.com/MiteshSharma/project/auth/app"
	"github.com/MiteshSharma/project/auth/repository"
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
	ii.Bus.AddHandler(model.CREATE_AUTH, ii.createAuth)
}

func (ii *InternalInput) createAuth(msg interface{}) (interface{}, error) {
	user := msg.(model.User)
	if user == (model.User{}) {
		return nil, errors.New("User object received is nil")
	}
	if err := user.Valid(); err != nil {
		return nil, errors.New("User object received is invalid")
	}

	userAuth, err := ii.App.UserLogin(&user)
	if err != nil {
		return nil, errors.New("User auth failed")
	}

	return userAuth, nil
}
