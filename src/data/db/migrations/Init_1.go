package migrations

import (
	"github.com/sajadblnyn/autocar-apis/config"
	"github.com/sajadblnyn/autocar-apis/constants"
	"github.com/sajadblnyn/autocar-apis/data/db"
	"github.com/sajadblnyn/autocar-apis/data/models"
	"github.com/sajadblnyn/autocar-apis/pkg/logging"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var logger = logging.NewLogger(config.GetConfig())

func Up_1(d db.DbService) {
	database := d.GetDb()

	country := models.Country{}
	city := models.City{}
	user_role := models.UserRole{}
	role := models.Role{}
	user := models.User{}
	persianYear := models.PersianYear{}
	color := models.Color{}
	file := models.File{}
	gearbox := models.Gearbox{}
	carType := models.CarType{}
	company := models.Company{}
	propertyCategory := models.PropertyCategory{}
	property := models.Property{}
	carModel := models.CarModel{}
	carModelColor := models.CarModelColor{}
	carModelYear := models.CarModelYear{}
	carModelImage := models.CarModelImage{}
	carModelPriceHistory := models.CarModelPriceHistory{}
	carModelProperty := models.CarModelProperty{}
	carModelComment := models.CarModelComment{}

	createTablesIfNotExists(database, country, city, user, role, user_role,
		persianYear, color, file, gearbox, carType, company, propertyCategory,
		property, carModel, carModelColor, carModelYear, carModelImage,
		carModelPriceHistory, carModelProperty, carModelComment)

	createDefaultInformation(database)
	createCountry(database)

	logger.Info(logging.Database, logging.Migration, "initial migration (up_1) has done successfully", nil)
}

func createTablesIfNotExists(d *gorm.DB, models ...interface{}) {
	tables := []interface{}{}
	for i := 0; i < len(models); i++ {
		if !d.Migrator().HasTable(models[i]) {
			tables = append(tables, models[i])
		}
	}
	d.Migrator().CreateTable(tables...)

}

func createDefaultInformation(d *gorm.DB) {
	adminRole := models.Role{Name: constants.AdminRoleName}
	createRoleIfNotExists(d, &adminRole)
	createRoleIfNotExists(d, &models.Role{Name: constants.DefaultRoleName})

	pass, _ := bcrypt.GenerateFromPassword([]byte(config.GetConfig().Server.DefaultAdminPassword), bcrypt.DefaultCost)
	user := models.User{
		Username: constants.DefaultUserName, Password: string(pass),
		FirstName: "test", LastName: "test",
		MobileNumber: "09369919691", Email: "sajadbalaniyan.sb@gmail.com"}

	createAdminUserIfNotExists(d, &user, adminRole.Id)

}

func createRoleIfNotExists(d *gorm.DB, adminRole *models.Role) {
	exists := 0
	d.Model(&models.Role{}).Select("1").Where("name=?", adminRole.Name).First(&exists)

	if exists == 0 {
		d.Create(adminRole)
	}
}

func createAdminUserIfNotExists(d *gorm.DB, u *models.User, roleId int) {
	exists := 0
	d.Model(&models.User{}).Select("1").Where("username=?", u.Username).First(&exists)

	if exists == 0 {
		d.Create(u)
		ur := models.UserRole{UserId: u.Id, RoleId: roleId}
		d.Create(&ur)
	}
}

func createCountry(database *gorm.DB) {
	count := 0
	database.
		Model(&models.Country{}).
		Select("count(*)").
		Find(&count)
	if count == 0 {
		database.Create(&models.Country{Name: "Iran", Cities: []models.City{
			{Name: "Tehran"},
			{Name: "Isfahan"},
			{Name: "Shiraz"},
			{Name: "Chalus"},
			{Name: "Ahwaz"},
		}})
		database.Create(&models.Country{Name: "USA", Cities: []models.City{
			{Name: "New York"},
			{Name: "Washington"},
		}})
		database.Create(&models.Country{Name: "Germany", Cities: []models.City{
			{Name: "Berlin"},
			{Name: "Munich"},
		}})
		database.Create(&models.Country{Name: "China", Cities: []models.City{
			{Name: "Beijing"},
			{Name: "Shanghai"},
		}})
		database.Create(&models.Country{Name: "Italy", Cities: []models.City{
			{Name: "Roma"},
			{Name: "Turin"},
		}})
		database.Create(&models.Country{Name: "France", Cities: []models.City{
			{Name: "Paris"},
			{Name: "Lyon"},
		}})
		database.Create(&models.Country{Name: "Japan", Cities: []models.City{
			{Name: "Tokyo"},
			{Name: "Kyoto"},
		}})
		database.Create(&models.Country{Name: "South Korea", Cities: []models.City{
			{Name: "Seoul"},
			{Name: "Ulsan"},
		}})
	}
}

func Down_1() {

}
