package cache

import (
	"sync"
	"time"
)

type Cache[T any] interface {
	Get(key string) (T, bool)
	Set(key string, value T)
}

type Entry[T any] struct {
	Value     T
	Timestamp time.Time
}

type TTLCache[T any] struct {
	expiration time.Duration
	data       map[string]Entry[T]
	mu         sync.RWMutex
}

func NewTTLCache[T any](expiration time.Duration) *TTLCache[T] {
	cache := &TTLCache[T]{
		expiration: expiration,
		data:       make(map[string]Entry[T]),
		mu:         sync.RWMutex{},
	}
	go cache.cleanup()
	return cache
}

func (c *TTLCache[T]) cleanup() {
	ticker := time.NewTicker(3 * time.Minute)
	for range ticker.C {
		c.mu.Lock()
		for key, entry := range c.data {
			if time.Since(entry.Timestamp) > c.expiration {
				delete(c.data, key)
			}
		}
		c.mu.Unlock()
	}
}

func (c *TTLCache[T]) Get(key string) (T, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	entry, ok := c.data[key]
	if !ok || time.Since(entry.Timestamp) > c.expiration {
		return entry.Value, false
	}
	return entry.Value, true
}

func (c *TTLCache[T]) Set(key string, value T) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = Entry[T]{
		Timestamp: time.Now(),
		Value:     value,
	}
}
