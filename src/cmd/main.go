package main

import (
	"log"

	"github.com/sajadblnyn/autocar-apis/apis"
	"github.com/sajadblnyn/autocar-apis/config"
	"github.com/sajadblnyn/autocar-apis/data/cache"
	"github.com/sajadblnyn/autocar-apis/data/db"
)

func main() {
	cfg := config.GetConfig()

	cacheService := cache.New()
	err := cacheService.Init(cfg)
	defer cacheService.Close()
	if err != nil {
		log.Fatal(err.Error())
	}

	dbService := db.New()
	err = dbService.Init(cfg)
	defer dbService.Close()
	if err != nil {
		log.Fatal(err.Error())
	}

	apis.InitServer(cfg)

}
