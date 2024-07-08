package cache

import (
	"sync"
	"time"

	"github.com/sajadblnyn/autocar-apis/config"
)

var once sync.Once
var cacheService CacheService

type CacheService interface {
	Init(cfg *config.Config) error
	Get(key string, value interface{}) error
	Set(key string, value interface{}, secondsExpireTimeDuration time.Duration) error
	Close()
}

func New() CacheService {
	once.Do(func() {
		cacheService = newRedis()

	})
	return cacheService
}
