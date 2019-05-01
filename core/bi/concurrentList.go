package bi

import (
	"encoding/json"
	"reflect"
	"sync"
)

type ConcurrentList struct {
	Items []map[string]interface{}
	sync.RWMutex
}

func NewConcurrentList() *ConcurrentList {
	return &ConcurrentList{Items: make([]map[string]interface{}, 0)}
}

func (c ConcurrentList) Length() int {
	c.Lock()
	defer c.Unlock()
	return len(c.Items)
}

func (c *ConcurrentList) Add(value map[string]interface{}) {
	c.Lock()
	defer c.Unlock()

	c.Items = append(c.Items, value)
}

func (c *ConcurrentList) Remove(value map[string]interface{}) {
	c.Lock()
	defer c.Unlock()

	idx := c.indexOf(value)
	if idx == -1 {
		return
	}

	c.Items = append(c.Items[:idx], c.Items[idx+1:]...)
}

func (c ConcurrentList) indexOf(value map[string]interface{}) int {
	if len(c.Items) == 0 {
		return -1
	}

	idx := -1
	for index := 0; index < len(c.Items); index++ {
		if reflect.DeepEqual(value, c.Items[index]) {
			idx = index
		}
	}
	return idx
}

func (c ConcurrentList) ToJSON() []byte {
	c.RLock()
	defer c.RUnlock()

	if jsonResponse, err := json.Marshal(c); err != nil {
		return nil
	} else {
		return jsonResponse
	}
}

func (c *ConcurrentList) GetAllAndFlush() []map[string]interface{} {
	c.Lock()
	defer c.Unlock()

	if len(c.Items) != 0 {
		items := make([]map[string]interface{}, len(c.Items))
		copy(items, c.Items)
		c.Items = make([]map[string]interface{}, 0)
		return items
	}
	return nil
}
