package bi

import (
	"time"

	"github.com/MiteshSharma/project/core/bi/flush"
	"github.com/MiteshSharma/project/core/logger"

	"github.com/MiteshSharma/project/core/bus"
)

const FLUSH_PERIOD_SEC int = 5
const FLUSH_LIST_SIZE int = 60

type BiReceiverHandler struct {
	List    *ConcurrentList
	Flusher flush.FlushEvent
	Log     logger.Logger
}

func NewBiReceiverHandler(logger logger.Logger, bus bus.Bus) *BiReceiverHandler {
	biReceiverHandler := &BiReceiverHandler{
		List:    NewConcurrentList(),
		Flusher: flush.NewBiModuleFlushEvent(bus),
		Log:     logger,
	}
	return biReceiverHandler
}

func (bi BiReceiverHandler) Init(bus bus.Bus) {
	bi.handlePeriodicFlush()
	bus.AddHandler("bi", bi.handleBiEvent)
}

func (bi BiReceiverHandler) handlePeriodicFlush() {
	go func() {
		for {
			<-time.After(time.Duration(FLUSH_PERIOD_SEC) * time.Second)
			go bi.onMessageReceive(nil, true)
		}
	}()
}

func (bi BiReceiverHandler) handleBiEvent(msg interface{}) (interface{}, error) {
	bi.onMessageReceive(msg.(map[string]interface{}), false)
	return nil, nil
}

func (bi BiReceiverHandler) onMessageReceive(value map[string]interface{}, isFlush bool) {
	if value != nil {
		bi.List.Add(value)
		if bi.List.Length() >= FLUSH_LIST_SIZE {
			isFlush = true
		}
	}
	if isFlush {
		bi.flushBiEvents()
	}
}

func (bi BiReceiverHandler) flushBiEvents() {
	itemsToFlush := bi.List.GetAllAndFlush()
	if itemsToFlush == nil {
		return
	}
	bi.Log.Info("BIEvent: Flushing BI events.", logger.Int("size", len(itemsToFlush)))
	bi.Flusher.Flush(itemsToFlush)
}
