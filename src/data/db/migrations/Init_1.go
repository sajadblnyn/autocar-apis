package migrations

import (
	"github.com/sajadblnyn/autocar-apis/config"
	"github.com/sajadblnyn/autocar-apis/data/db"
	"github.com/sajadblnyn/autocar-apis/data/models"
	"github.com/sajadblnyn/autocar-apis/pkg/logging"
)

var logger = logging.NewLogger(config.GetConfig())

func Up_1(d db.DbService) {
	database := d.GetDb()

	country := models.Country{}
	city := models.City{}

	tables := []interface{}{}
	if !database.Migrator().HasTable(country) {
		tables = append(tables, country)
	}

	if !database.Migrator().HasTable(city) {
		tables = append(tables, city)
	}

	database.Migrator().CreateTable(tables...)
	logger.Info(logging.Database, logging.Migration, "initial migration (up_1) has done successfully", nil)
}

func Down_1() {

}
