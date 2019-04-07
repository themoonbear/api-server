package utils

import (
	"github.com/patrickmn/go-cache"
)

var defaultCache *cache.Cache

// AddCache 缓存数据
func AddCache(key, data string) {
	if defaultCache != nil {
		defaultCache.Set(key, data, cache.NoExpiration)
	}
}

// GetCache 从缓存获取数据
func GetCache(key string) (string, bool) {
	if defaultCache == nil {
		return "", false
	}
	data, ok := defaultCache.Get(key)
	if !ok {
		return "", false
	}

	r, ok := data.(string)
	if !ok || r == "" {
		return "", false
	}

	return r, true
}
