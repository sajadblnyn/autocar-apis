package cache

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/sajadblnyn/autocar-apis/config"
	"github.com/sajadblnyn/autocar-apis/pkg/logging"
)

var logger = logging.NewLogger(config.GetConfig())

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
	logger.Info(logging.CacheService, logging.Startup, "cache service connection established successfully", nil)

	return nil

}

func (r *redisService) Close() {
	r.redisClient.Close()
}

func (r *redisService) Set(key string, value interface{}, expireTimeDuration time.Duration) error {
	v, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return r.redisClient.Set(key, v, expireTimeDuration).Err()
}

func (r *redisService) Get(key string, value interface{}) error {

	v, err := r.redisClient.Get(key).Result()

	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(v), value)
	if err != nil {
		return err
	}
	return nil
}
