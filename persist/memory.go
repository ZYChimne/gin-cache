package persist

import (
	"reflect"
	"time"

	"github.com/jellydator/ttlcache/v3"
)

// MemoryStore local memory cache store
type MemoryStore struct {
	Cache *ttlcache.Cache[string, interface{}]
}

// NewMemoryStore allocate a local memory store with default expiration
func NewMemoryStore(defaultExpiration time.Duration) *MemoryStore {
	cacheStore := ttlcache.New[string, interface{}](
		ttlcache.WithTTL[string, interface{}](defaultExpiration),
		ttlcache.WithDisableTouchOnHit[string, interface{}](),
	)
	return &MemoryStore{
		Cache: cacheStore,
	}
}

// Set put key value pair to memory store, and expire after expireDuration
func (c *MemoryStore) Set(key string, value interface{}, expireDuration time.Duration) error {
	c.Cache.DeleteExpired()
	c.Cache.Set(key, value, expireDuration)
	return nil
}

// Delete remove key in memory store, do nothing if key doesn't exist
func (c *MemoryStore) Delete(key string) error {
	if !c.Cache.Has(key) {
		return ErrCacheMiss
	}
	c.Cache.Delete(key)
	return nil
}

// Get get key in memory store, if key doesn't exist, return ErrCacheMiss
func (c *MemoryStore) Get(key string, value interface{}) error {
	c.Cache.DeleteExpired()
	if !c.Cache.Has(key) {
		return ErrCacheMiss
	}
	val := c.Cache.Get(key)
	v := reflect.ValueOf(value)
	v.Elem().Set(reflect.ValueOf(val.Value()))
	return nil
}
