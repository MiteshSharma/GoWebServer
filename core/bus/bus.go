package bus

import (
	"reflect"

	"github.com/MiteshSharma/project/core/logger"
)

type HandlerFunc = func(msg interface{}) (interface{}, error)

type Bus interface {
	Publish(messageType string, msg interface{}) (interface{}, error)
	AddHandler(messageType string, handler HandlerFunc) error
}

type AppBus struct {
	Log       logger.Logger
	listeners map[string][]HandlerFunc
}

func NewBus(logger logger.Logger) Bus {
	bus := &AppBus{
		Log: logger,
	}
	bus.listeners = make(map[string][]HandlerFunc)
	return bus
}

func (b *AppBus) Publish(messageType string, msg interface{}) (interface{}, error) {
	b.Log.Debug("Bus: Message received by bus with type", logger.String("type", messageType))
	_, isExists := b.listeners[messageType]
	if isExists {
		for _, listenerHandler := range b.listeners[messageType] {
			var params = make([]reflect.Value, 1)
			params[0] = reflect.ValueOf(msg)

			ret := reflect.ValueOf(listenerHandler).Call(params)
			response := ret[0].Interface()
			err := ret[1].Interface()
			if err == nil {
				return response, nil
			}
			return response, err.(error)
		}
	}
	return nil, nil
}

func (b *AppBus) AddHandler(messageType string, handler HandlerFunc) error {
	b.Log.Debug("Message handler added to bus with type : " + messageType)
	_, isExists := b.listeners[messageType]
	if !isExists {
		b.listeners[messageType] = make([]HandlerFunc, 0)
	}
	b.listeners[messageType] = append(b.listeners[messageType], handler)
	return nil
}
