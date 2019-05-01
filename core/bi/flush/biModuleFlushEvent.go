package flush

import (
	"github.com/MiteshSharma/project/core/bus"
)

type BiModuleFlushEvent struct {
	Bus bus.Bus
}

func NewBiModuleFlushEvent(bus bus.Bus) *BiModuleFlushEvent {
	flusher := &BiModuleFlushEvent{
		Bus: bus,
	}
	return flusher
}

func (f BiModuleFlushEvent) Flush(events []map[string]interface{}) error {
	f.Bus.Publish("SEND_BI_BATCH_EVENTS", events)
	return nil
}
