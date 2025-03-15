package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createAt time.Time
	val      []byte
}

type Cache struct {
	entries  map[string]cacheEntry
	mu       sync.Mutex
	interval time.Duration
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		entries:  make(map[string]cacheEntry),
		mu:       sync.Mutex{},
		interval: interval,
	}

	go c.reapLoop()

	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.entries[key] = cacheEntry{
		createAt: time.Now(),
		val:      val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.entries[key]
	if !ok {
		return nil, false
	}

	return entry.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for range ticker.C {
		c.reap()
	}
}

func (c *Cache) reap() {
	c.mu.Lock()
	defer c.mu.Unlock()

	now := time.Now()
	for k, v := range c.entries {
		if now.Sub(v.createAt) > c.interval {
			delete(c.entries, k)
		}
	}
}
