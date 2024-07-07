package db

import (
	"fmt"
	"time"

	"github.com/sajadblnyn/autocar-apis/config"
	"github.com/sajadblnyn/autocar-apis/pkg/logging"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresDb struct {
	dbClient *gorm.DB
}

var logger = logging.NewLogger(config.GetConfig())

func (p *postgresDb) Init(cfg *config.Config) error {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=Asia/Tehran",
		cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.DbName, cfg.Postgres.SSLMode)
	var err error

	p.dbClient, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return err
	}

	sqlDb, _ := p.dbClient.DB()
	err = sqlDb.Ping()
	if err != nil {
		return err
	}

	sqlDb.SetMaxIdleConns(cfg.Postgres.MaxIdleConns)
	sqlDb.SetMaxOpenConns(cfg.Postgres.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(cfg.Postgres.ConnMaxLifetime * time.Minute)

	logger.Info(logging.Database, logging.Startup, "database connection established successfully", nil)
	return nil
}
func (p *postgresDb) Close() {
	conn, _ := p.dbClient.DB()
	conn.Close()
}

func (p *postgresDb) GetDb() *gorm.DB {
	return p.dbClient
}

func newPostgresDb() *postgresDb {
	return &postgresDb{}
}
