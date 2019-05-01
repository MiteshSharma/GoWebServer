package bus

type TestBus struct {
}

func NewTestBus() Bus {
	bus := &TestBus{}
	return bus
}

func (b *TestBus) Publish(messageType string, msg interface{}) (interface{}, error) {
	return nil, nil
}

func (b *TestBus) AddHandler(messageType string, handler HandlerFunc) error {
	return nil
}
