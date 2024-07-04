package cache

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/sajadblnyn/autocar-apis/config"
)

type Redis struct {
	redisClient *redis.Client
}

func newRedis() *Redis {
	return &Redis{}
}
func (r *Redis) Init(cfg *config.Config) {
	r.redisClient = redis.NewClient(&redis.Options{
		Addr:               fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port),
		Password:           cfg.Redis.Password,
		DB:                 0,
		DialTimeout:        time.Second * cfg.Redis.DialTimeout,
		ReadTimeout:        time.Second * cfg.Redis.ReadTimeout,
		WriteTimeout:       time.Second * cfg.Redis.WriteTimeout,
		PoolSize:           cfg.Redis.PoolSize,
		PoolTimeout:        cfg.Redis.PoolTimeout,
		IdleTimeout:        500 * time.Microsecond,
		IdleCheckFrequency: cfg.Redis.IdleCheckFrequency * time.Microsecond,
	})

}

func (r *Redis) Close() {
	r.redisClient.Close()
}
