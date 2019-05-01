package app

import (
	"encoding/json"

	"github.com/MiteshSharma/project/core/logger"
)

func (a *App) HandleBIBatch(events []map[string]interface{}) {
	mqEvent := map[string]interface{}{
		"type":  "batch",
		"event": events,
	}
	msg, err := json.Marshal(mqEvent)
	if err != nil {
		a.Log.Debug("Error marshalling events batch", logger.Error(err))
	}
	err = a.MqAgent.Write(msg)
	if err != nil {
		a.Log.Debug("sending message to MQ failed", logger.Error(err))
	}
}

func (a *App) HandleBIEvent(event map[string]interface{}) {
	mqEvent := map[string]interface{}{
		"type":  "single",
		"event": event,
	}
	msg, err := json.Marshal(mqEvent)
	if err != nil {
		a.Log.Debug("Error marshalling event", logger.Error(err))
	}
	err = a.MqAgent.Write(msg)
	if err != nil {
		a.Log.Debug("sending message to MQ failed", logger.Error(err))
	}
}
