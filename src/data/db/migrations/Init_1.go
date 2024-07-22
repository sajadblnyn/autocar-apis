package migrations

import (
	"time"

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

	createDefaultUserInformation(database)
	createPropertyCategory(database)

	createCountry(database)
	createCarType(database)
	createGearbox(database)
	createColor(database)
	createYear(database)
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

func createDefaultUserInformation(d *gorm.DB) {
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
func createCarType(database *gorm.DB) {
	count := 0
	database.
		Model(&models.CarType{}).
		Select("count(*)").
		Find(&count)
	if count == 0 {
		database.Create(&models.CarType{Name: "Crossover"})
		database.Create(&models.CarType{Name: "Sedan"})
		database.Create(&models.CarType{Name: "Sports"})
		database.Create(&models.CarType{Name: "Coupe"})
		database.Create(&models.CarType{Name: "Hatchback"})
	}
}
func createColor(database *gorm.DB) {
	count := 0
	database.
		Model(&models.Color{}).
		Select("count(*)").
		Find(&count)
	if count == 0 {
		database.Create(&models.Color{Name: "Black", HexCode: "#000000"})
		database.Create(&models.Color{Name: "White", HexCode: "#ffffff"})
		database.Create(&models.Color{Name: "Blue", HexCode: "#0000ff"})
	}
}

func createYear(database *gorm.DB) {
	count := 0
	database.
		Model(&models.PersianYear{}).
		Select("count(*)").
		Find(&count)
	if count == 0 {

		database.Create(&models.PersianYear{
			PersianTitle: "1402",
			Year:         1402,
			StartAt:      time.Date(2023, time.Month(3), 21, 0, 0, 0, 0, time.UTC),
			EndAt:        time.Date(2024, time.Month(3), 20, 0, 0, 0, 0, time.UTC),
		})

		database.Create(&models.PersianYear{
			PersianTitle: "1401",
			Year:         1401,
			StartAt:      time.Date(2022, time.Month(3), 21, 0, 0, 0, 0, time.UTC),
			EndAt:        time.Date(2023, time.Month(3), 21, 0, 0, 0, 0, time.UTC),
		})

		database.Create(&models.PersianYear{
			PersianTitle: "1400",
			Year:         1400,
			StartAt:      time.Date(2021, time.Month(3), 21, 0, 0, 0, 0, time.UTC),
			EndAt:        time.Date(2022, time.Month(3), 21, 0, 0, 0, 0, time.UTC),
		})

		database.Create(&models.PersianYear{
			PersianTitle: "1399",
			Year:         1399,
			StartAt:      time.Date(2020, time.Month(3), 20, 0, 0, 0, 0, time.UTC),
			EndAt:        time.Date(2021, time.Month(3), 21, 0, 0, 0, 0, time.UTC),
		})

		database.Create(&models.PersianYear{
			PersianTitle: "1398",
			Year:         1398,
			StartAt:      time.Date(2019, time.Month(3), 21, 0, 0, 0, 0, time.UTC),
			EndAt:        time.Date(2020, time.Month(3), 20, 0, 0, 0, 0, time.UTC),
		})

		database.Create(&models.PersianYear{
			PersianTitle: "1398",
			Year:         1398,
			StartAt:      time.Date(2018, time.Month(3), 21, 0, 0, 0, 0, time.UTC),
			EndAt:        time.Date(2019, time.Month(3), 21, 0, 0, 0, 0, time.UTC),
		})
	}
}

func createGearbox(database *gorm.DB) {
	count := 0
	database.
		Model(&models.Gearbox{}).
		Select("count(*)").
		Find(&count)
	if count == 0 {
		database.Create(&models.Gearbox{Name: "Manual"})
		database.Create(&models.Gearbox{Name: "Automatic"})
	}
}

func createPropertyCategory(database *gorm.DB) {
	count := 0

	database.
		Model(&models.PropertyCategory{}).
		Select("count(*)").
		Find(&count)
	if count == 0 {
		database.Create(&models.PropertyCategory{Name: "Body"})                     // بدنه
		database.Create(&models.PropertyCategory{Name: "Engine"})                   // موتور
		database.Create(&models.PropertyCategory{Name: "Drivetrain"})               // پیشرانه
		database.Create(&models.PropertyCategory{Name: "Suspension"})               // تعلیق
		database.Create(&models.PropertyCategory{Name: "Equipment"})                // تجهیزات
		database.Create(&models.PropertyCategory{Name: "Driver support systems"})   // سیستم های پشتیبانی راننده
		database.Create(&models.PropertyCategory{Name: "Lights"})                   // چراغ ها
		database.Create(&models.PropertyCategory{Name: "Multimedia"})               // چند رسانه ای
		database.Create(&models.PropertyCategory{Name: "Safety equipment"})         // تجهیزات ایمنی
		database.Create(&models.PropertyCategory{Name: "Seats and steering wheel"}) // صندلی و فرمان
		database.Create(&models.PropertyCategory{Name: "Windows and mirrors"})      // پنجره و آینه
	}
	createProperty(database, "Body")
	createProperty(database, "Engine")
	createProperty(database, "Drivetrain")
	createProperty(database, "Suspension")
	createProperty(database, "Comfort")
	createProperty(database, "Driver support systems")
	createProperty(database, "Lights")
	createProperty(database, "Multimedia")
	createProperty(database, "Safety equipment")
	createProperty(database, "Seats and steering wheel")
	createProperty(database, "Windows and mirrors")

}
func createProperty(database *gorm.DB, cat string) {
	count := 0
	catModel := models.PropertyCategory{}

	database.
		Model(models.PropertyCategory{}).
		Where("name = ?", cat).
		Find(&catModel)

	database.
		Model(&models.Property{}).
		Select("count(*)").
		Where("category_id = ?", catModel.Id).
		Find(&count)

	if count > 0 || catModel.Id == 0 {
		return
	}
	var props *[]models.Property
	switch cat {
	case "Body":
		props = getBodyProperties(catModel.Id)

	case "Engine":
		props = getEngineProperties(catModel.Id)

	case "Drivetrain":
		props = getDrivetrainProperties(catModel.Id)

	case "Suspension":
		props = getSuspensionProperties(catModel.Id)

	case "Comfort":
		props = getComfortProperties(catModel.Id)

	case "Driver support systems":
		props = getDriverSupportSystemProperties(catModel.Id)

	case "Lights":
		props = getLightsProperties(catModel.Id)

	case "Multimedia":
		props = getMultimediaProperties(catModel.Id)

	case "Safety equipment":
		props = getSafetyEquipmentProperties(catModel.Id)

	case "Seats and steering wheel":
		props = getSeatsProperties(catModel.Id)

	case "Windows and mirrors":
		props = getWindowsProperties(catModel.Id)

	default:
		props = &([]models.Property{})
	}

	for _, prop := range *props {
		database.Create(&prop)
	}
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
		}, Companies: []models.Company{
			{Name: "Saipa"},
			{Name: "Iran khodro"},
		}})
		database.Create(&models.Country{Name: "USA", Cities: []models.City{
			{Name: "New York"},
			{Name: "Washington"},
		}, Companies: []models.Company{
			{Name: "Tesla"},
			{Name: "Jeep"},
		}})
		database.Create(&models.Country{Name: "Germany", Cities: []models.City{
			{Name: "Berlin"},
			{Name: "Munich"},
		}, Companies: []models.Company{
			{Name: "Opel"},
			{Name: "Benz"},
		}})
		database.Create(&models.Country{Name: "China", Cities: []models.City{
			{Name: "Beijing"},
			{Name: "Shanghai"},
		}, Companies: []models.Company{
			{Name: "Chery"},
			{Name: "Geely"},
		}})
		database.Create(&models.Country{Name: "Italy", Cities: []models.City{
			{Name: "Roma"},
			{Name: "Turin"},
		}, Companies: []models.Company{
			{Name: "Ferrari"},
			{Name: "Fiat"},
		}})
		database.Create(&models.Country{Name: "France", Cities: []models.City{
			{Name: "Paris"},
			{Name: "Lyon"},
		}, Companies: []models.Company{
			{Name: "Renault"},
			{Name: "Bugatti"},
		}})
		database.Create(&models.Country{Name: "Japan", Cities: []models.City{
			{Name: "Tokyo"},
			{Name: "Kyoto"},
		}, Companies: []models.Company{
			{Name: "Toyota"},
			{Name: "Honda"},
		}})
		database.Create(&models.Country{Name: "South Korea", Cities: []models.City{
			{Name: "Seoul"},
			{Name: "Ulsan"},
		}, Companies: []models.Company{
			{Name: "Kia"},
			{Name: "Hyundai"},
		}})
	}
}

func Down_1() {

}
