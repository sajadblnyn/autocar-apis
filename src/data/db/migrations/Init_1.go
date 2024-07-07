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
	createTablesIfNotExists(database, city, country, user, role, user_role)

	createDefaultInformation(database)

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

func Down_1() {

}
