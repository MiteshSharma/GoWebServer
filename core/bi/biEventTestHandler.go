package bi

import (
	"errors"
)

// BiEvent struct to send all BI events
type BiTestEventHandler struct {
}

func NewBiTestEventHandler() *BiTestEventHandler {
	biTestEventHandler := &BiTestEventHandler{}
	return biTestEventHandler
}

func (bi BiTestEventHandler) Send(eventName string, data map[string]interface{}) error {
	if eventName == "" {
		return errors.New("event name must be provided")
	}
	return nil
}
