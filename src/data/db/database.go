package db

import "github.com/sajadblnyn/autocar-apis/config"

type DbService interface {
	Init(cfg *config.Config) error
	Close()
}

func New() DbService {
	return newPostgresDb()
}
