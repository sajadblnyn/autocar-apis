package main

import (
	"github.com/sajadblnyn/autocar-apis/apis"
	"github.com/sajadblnyn/autocar-apis/config"
	"github.com/sajadblnyn/autocar-apis/data/cache"
	"github.com/sajadblnyn/autocar-apis/data/db"
	"github.com/sajadblnyn/autocar-apis/pkg/logging"
)

func main() {
	cfg := config.GetConfig()

	logger := logging.NewLogger(cfg)

	cacheService := cache.New()
	err := cacheService.Init(cfg)
	defer cacheService.Close()
	if err != nil {
		logger.Fatal(logging.CacheService, logging.Startup, err.Error(), nil)
	}

	dbService := db.New()
	err = dbService.Init(cfg)
	defer dbService.Close()
	if err != nil {
		logger.Fatal(logging.Database, logging.Startup, err.Error(), nil)
	}

	apis.InitServer(cfg)

}
