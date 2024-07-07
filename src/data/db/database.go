package db

import (
	"github.com/sajadblnyn/autocar-apis/config"
	"gorm.io/gorm"
)

type DbService interface {
	Init(cfg *config.Config) error
	Close()
	GetDb() *gorm.DB
}

func New() DbService {
	return newPostgresDb()
}
