package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	CreatedAt time.Time
	Val       []byte
}

type cache struct {
	mutex *sync.Mutex
	Data  map[string]cacheEntry
}

func (c *cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	cacheEntry := cacheEntry{
		CreatedAt: time.Now().UTC(),
		Val:       val,
	}
	c.Data[key] = cacheEntry

}

func (c *cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	value, exists := c.Data[key]
	return value.Val, exists

}

func (c *cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *cache) reap(now time.Time, last time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	for k, v := range c.Data {
		if v.CreatedAt.Before(now.Add(-last)) {
			delete(c.Data, k)
		}
	}
}

func Newcache(interval time.Duration) cache {
	newCache := cache{
		Data:  make(map[string]cacheEntry),
		mutex: &sync.Mutex{},
	}
	go newCache.reapLoop(interval)
	return newCache
}
