package main

import (
	"github.com/sajadblnyn/autocar-apis/apis"
	"github.com/sajadblnyn/autocar-apis/config"
	"github.com/sajadblnyn/autocar-apis/data/cache"
)

func main() {
	cfg := config.GetConfig()
	apis.InitServer(cfg)

	cacheService := cache.New()
	cacheService.Init(cfg)
	defer cacheService.Close()

}
