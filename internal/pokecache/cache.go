package pokecache

import (
	"fmt"
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

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
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
		createdAt: time.Now(),
		val:       val,
	}

	c.entries[key] = entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mtx.RLock()
	defer c.mtx.RUnlock()

	entry, ok := c.entries[key]
	if !ok {
		return nil, false
	}

	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			toRemove := make([]string, 0)
			c.mtx.Lock()
			for key, entry := range c.entries {
				if entry.createdAt.Add(interval).Before(time.Now()) {
					fmt.Println("Add to list")
					toRemove = append(toRemove, key)
				}
			}
			for _, key := range toRemove {
				delete(c.entries, key)
			}
			c.mtx.Unlock()
		}
	}
}
