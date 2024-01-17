package pokecache

import (
	"time"
	"sync"
)

type Cache struct {
	entries map[string]cacheEntry
	interval time.Duration
	mut sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

// NewCache creates a new instance of Cache
func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		entries:  make(map[string]cacheEntry),
		interval: interval,
		mut:      sync.Mutex{},
	}
	go c.reapLoop()
	return c
}

// Add a new entry to the cache
func (c *Cache) Add(key string, val []byte) {
	c.mut.Lock()
	defer c.mut.Unlock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

// Get an entry from the cache.
// Returns the data and a boolean indicating whether the data was found or not
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mut.Lock()
	defer c.mut.Unlock()
	entry, found := c.entries[key]
	if !found {
		return nil, false
	}
	return entry.val, true
}

// reapLoop scans the cache for old entries and removes them
func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for range ticker.C {
		c.mut.Lock()
		for key, entry := range c.entries {
			if time.Since(entry.createdAt) > c.interval {
				delete(c.entries, key)
			}
		}
		c.mut.Unlock()
	}
}
