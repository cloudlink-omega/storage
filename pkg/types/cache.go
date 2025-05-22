package types

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/patrickmn/go-cache"
)

// DBCache is a basic key-value in-memory cache for the database.
type DBCache struct {
	cache *cache.Cache
	lock  *sync.RWMutex
}

func NewDBCache() *DBCache {
	c := &DBCache{}
	c.init()
	return c
}

func (c *DBCache) init() {
	if c.cache == nil {
		c.cache = cache.New(5*time.Minute, 2*time.Minute)
		c.lock = &sync.RWMutex{}
	}
}

func (c *DBCache) make_key(keytype string, keys ...string) string {
	return fmt.Sprintf("%s_%s", keytype, strings.Join(keys, ";"))
}

func (c *DBCache) Set(keytype string, value any, keys ...string) {
	c.init()
	key := c.make_key(keytype, keys...)
	c.lock.Lock()
	defer c.lock.Unlock()
	c.cache.SetDefault(key, value)
}

func (c *DBCache) Get(keytype string, keys ...string) (any, bool) {
	c.init()
	key := c.make_key(keytype, keys...)

	c.lock.RLock()
	defer c.lock.RUnlock()
	data, hit := c.cache.Get(key)
	if hit {
		return data, true
	}
	return nil, false
}

func (c *DBCache) Delete(keytype string, keys ...string) {
	c.init()
	key := c.make_key(keytype, keys...)

	c.lock.Lock()
	defer c.lock.Unlock()
	c.cache.Delete(key)
}

func (c *DBCache) Flush() {
	c.init()
	c.lock.Lock()
	defer c.lock.Unlock()
	c.cache.Flush()
}
