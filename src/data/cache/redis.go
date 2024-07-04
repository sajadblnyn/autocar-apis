package cache

import (
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/sajadblnyn/autocar-apis/config"
)

type redisService struct {
	redisClient *redis.Client
}

func newRedis() *redisService {
	return &redisService{}
}
func (r *redisService) Init(cfg *config.Config) error {
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

	_, err := r.redisClient.Ping().Result()
	if err != nil {
		return err
	}
	log.Println("cache service connection established")

	return nil

}

func (r *redisService) Close() {
	r.redisClient.Close()
}
