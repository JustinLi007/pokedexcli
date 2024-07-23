package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries map[string]cacheEntry
	mtx     *sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		entries: make(map[string]cacheEntry),
		mtx:     &sync.RWMutex{},
	}

	go cache.reapLoop(interval)

	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	entry := cacheEntry{
		createdAt: time.Now().UTC(),
		val:       val,
	}

	c.entries[key] = entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mtx.RLock()
	defer c.mtx.RUnlock()

	entry, ok := c.entries[key]
	return entry.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			c.reap(time.Now().UTC(), interval)
		}
	}
}

func (c *Cache) reap(now time.Time, interval time.Duration) {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	for key, entry := range c.entries {
		if entry.createdAt.Before(now.Add(-interval)) {
			delete(c.entries, key)
		}
	}
}
