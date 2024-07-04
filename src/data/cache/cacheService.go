package cache

import (
	"github.com/sajadblnyn/autocar-apis/config"
)

type CacheService interface {
	Init(cfg *config.Config) error
	Close()
}

func New() CacheService {
	return newRedis()
}