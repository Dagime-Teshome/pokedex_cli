package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	CreatedAt time.Time
	Val       []byte
}

type Cache struct {
	mutex *sync.Mutex
	Data  map[string]cacheEntry
}
