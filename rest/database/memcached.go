package database

import (
	"github.com/bradfitz/gomemcache/memcache"
)

var Mem *memcache.Client

// Set permits to set a temporary value, on the cache
// via Memcached
func Set(key string, value string, time int32) {
	Mem.Set(&memcache.Item{
		Key:        key,
		Value:      []byte(value),
		Expiration: time,
	})
}
