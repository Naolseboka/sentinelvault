package localcache

import (
	"sync"
	"time"
)

type entry struct {
	value any
	exp   time.Time
}

type Cache struct {
	mu   sync.RWMutex
	data map[string]entry
}

func New() *Cache { return &Cache{data: map[string]entry{}} }

func (c *Cache) Set(key string, value any, ttl time.Duration) {
	c.mu.Lock()
	c.data[key] = entry{value: value, exp: time.Now().Add(ttl)}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) (any, bool) {
	c.mu.RLock()
	e, ok := c.data[key]
	c.mu.RUnlock()
	if !ok || time.Now().After(e.exp) {
		return nil, false
	}
	return e.value, true
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	delete(c.data, key)
	c.mu.Unlock()
}

func (c *Cache) GetOrSet(key string, ttl time.Duration, fn func() (any, error)) (any, error) {
	if v, ok := c.Get(key); ok {
		return v, nil
	}
	v, err := fn()
	if err != nil {
		return nil, err
	}
	c.Set(key, v, ttl)
	return v, nil
}

func (c *Cache) Len() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return len(c.data)
}

// maintenance note 34
