package cache

import (
	"github.com/patrickmn/go-cache"
)

var cacheO *cache.Cache

func getCache() *cache.Cache {

	if cacheO != nil {
		return cacheO
	}
	cacheO = cache.New(-1, 0)
	return cacheO
}
