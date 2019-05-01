package mq

import (
	"github.com/MiteshSharma/project/core/config"
	"github.com/MiteshSharma/project/core/errors"
)

type MqAgent interface {
	Read() (*MqObject, error)
	Write([]byte) error
	Delete(string) error
}

type MqObject struct {
	Id      string
	Message []byte
}

var agent MqAgent

func GetAgent(mqConfig config.MqConfig) MqAgent {
	if agent != nil {
		return agent
	} else {
		var err error
		switch mqConfig.Type {
		case "sqs":
			agent, err = NewSqsAgent(mqConfig)
		//case "kafka":
		// TODOs
		default:
			panic(&errors.MqTypeNotSupportedError{
				Type: mqConfig.Type,
			})
		}

		if err != nil {
			panic(err)
		}

		return agent
	}
}
