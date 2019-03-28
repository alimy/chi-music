package cache

import (
	"github.com/alimy/chi-music/pkg/json"
	"github.com/mediocregopher/radix/v3"
	"github.com/unisx/logus"

	"net/http"
)

// RedisCache Cache interface implement by redis
type RedisCache struct {
	// TODO
}

// EntryTo write cached entry to gin.Context
func (r *RedisCache) EntryTo(key string, w http.ResponseWriter) bool {
	content, err := r.entryFrom(key)
	if err == nil {
		w.WriteHeader(http.StatusOK)
		header := w.Header()
		header.Set("Content-Type", "application/json; charset=utf-8")
		_, err = w.Write(content)
	}
	if err != nil {
		return false
	}
	return true
}

// CacheFrom cache entry
func (r *RedisCache) CacheFrom(key string, entry interface{}) {
	jsonVal, err := json.Marshal(entry)
	if err != nil {
		logus.E("marshl entry", err)
		return
	}
	cache <- &jsonEntry{act: actJsonGet, key: key, val: string(jsonVal)}
}

func (r *RedisCache) Expire(key string) {
	cache <- &jsonEntry{act: actExpire, key: key}
}

func (r *RedisCache) entryFrom(key string) ([]byte, error) {
	var entry []byte
	err := rds.Do(radix.Cmd(&entry, "JSON.GET", key, "."))
	if err == nil && len(entry) == 0 {
		err = errNoExistEntry
	}
	return entry, err
}
