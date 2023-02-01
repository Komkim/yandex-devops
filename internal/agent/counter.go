package agent

import "sync"

type Counter struct {
	count int64
	mutex *sync.RWMutex
}

func NewCounter() *Counter {
	return &Counter{
		mutex: &sync.RWMutex{},
		count: 0,
	}
}

func (c *Counter) Inc() {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	c.count++
}

func (c *Counter) Dec() {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	c.count--
}

func (c *Counter) Reset() {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	c.count = 0
}

func (c *Counter) Get() int64 {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.count
}
