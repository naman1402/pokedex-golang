package pokecache

import (
	"sync"
	"time"
)

// main struct
type Cache struct {
	cache map[string]cacheEntry
	mu    *sync.Mutex
}

// representing single cache entry
type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
		mu:    &sync.Mutex{},
	}

	go c.reapLoop(interval)
	return c
}

func (c *Cache) AddEntry(key string, value []byte) {
	c.mu.Lock()

	defer c.mu.Unlock()
	// New entry in mapping
	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
}

func (c *Cache) GetEntry(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	// Mapping function in GO
	entry, exists := c.cache[key]
	return entry.val, exists
}

// ticker is used to generate a channel that sends time pulses after specific duration(interval)
// loop will be executed every time a new pulse is received
// reap function is used to remove expired entries from the cache, interval is expiration time
func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

// iterates over cache and check for expired entry
// deletes the expired entry and unlocks the c [defer]
func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for k, v := range c.cache {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.cache, k)
		}
	}
}
