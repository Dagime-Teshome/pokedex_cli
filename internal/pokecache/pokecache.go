package pokecache

import (
	"sync"
	"time"
)

func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	cacheEntry := cacheEntry{
		CreatedAt: time.Now().UTC(),
		Val:       val,
	}
	c.Data[key] = cacheEntry

}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	value, exists := c.Data[key]
	return value.Val, exists

}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	for k, v := range c.Data {
		if v.CreatedAt.Before(now.Add(-last)) {
			delete(c.Data, k)
		}
	}
}

func Newcache(interval time.Duration) *Cache {
	newCache := Cache{
		Data:  make(map[string]cacheEntry),
		mutex: &sync.Mutex{},
	}
	go newCache.reapLoop(interval)
	return &newCache
}
