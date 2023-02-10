package syncs

import "sync"

type Counter struct {
	value int
	mu sync.Mutex
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.value += 1
}

func (c *Counter) Value() int {
	return c.value
}
