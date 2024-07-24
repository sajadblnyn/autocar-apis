package apis

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/sajadblnyn/autocar-apis/apis/middlewares"
	"github.com/sajadblnyn/autocar-apis/apis/routers"
	"github.com/sajadblnyn/autocar-apis/apis/validations"
	"github.com/sajadblnyn/autocar-apis/config"
)

func InitServer(cfg *config.Config) {
	r := gin.New()

	RegisterCustomValidators()
	RegisterMiddlewares(r, cfg)
	RegisterRoutes(r, cfg)

	r.Run(fmt.Sprintf(":%s", cfg.Server.ExternalPort))
}

func RegisterCustomValidators() {
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		v.RegisterValidation("iran-mobile-validator", validations.IranianMobileValidator, true)
		v.RegisterValidation("password", validations.PasswordValidator, true)

	}
}
func RegisterRoutes(r *gin.Engine, cfg *config.Config) {
	v1 := r.Group("/api/v1/")
	{
		health := v1.Group("/health")
		routers.Health(health)

		test := v1.Group("/test")
		// test.Use(middlewares.TestMiddleware())
		routers.Test(test)

		users := v1.Group("/users")
		routers.User(users, cfg)

		country := v1.Group("/countries")
		country.Use(middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		routers.Country(country, cfg)

		city := v1.Group("/cities")
		city.Use(middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		routers.City(city, cfg)

		file := v1.Group("/files")
		file.Use(middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		routers.File(file, cfg)

		property := v1.Group("/properties")
		property.Use(middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		routers.Property(property, cfg)

		propertyCategory := v1.Group("/property-categories")
		property.Use(middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		routers.PropertyCategory(propertyCategory, cfg)

		carType := v1.Group("/car-types")
		carType.Use(middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		routers.CarType(carType, cfg)

		gearbox := v1.Group("/gearboxes")
		gearbox.Use(middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		routers.Gearbox(gearbox, cfg)

		company := v1.Group("/companies")
		company.Use(middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		routers.Company(company, cfg)

		carModel := v1.Group("/car-models")
		carModel.Use(middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		routers.CarModel(carModel, cfg)

		color := v1.Group("/colors")
		color.Use(middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		routers.Color(color, cfg)

		carModelColor := v1.Group("/car-model-colors")
		carModelColor.Use(middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		routers.CarModelColor(carModelColor, cfg)

		year := v1.Group("/persian-years")
		year.Use(middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		routers.PersianYear(year, cfg)

		carModelYear := v1.Group("/car-model-years")
		carModelYear.Use(middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		routers.CarModelPersianYear(carModelYear, cfg)

		carModelImage := v1.Group("/car-model-images")
		carModelImage.Use(middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		routers.CarModelImage(carModelImage, cfg)

		carModelProperty := v1.Group("/car-model-properties")
		carModelProperty.Use(middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		routers.CarModelProperty(carModelProperty, cfg)

		carModelYearPrHistory := v1.Group("/car-model-year-price-histories")
		carModelYearPrHistory.Use(middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		routers.CarModelPriceHistory(carModelYearPrHistory, cfg)

		r.Static("files", "../cmd/files")

	}
}

func RegisterMiddlewares(r *gin.Engine, cfg *config.Config) {
	r.Use(middlewares.Cors(cfg))
	r.Use(middlewares.DefaultStructuredLogger(cfg))
	r.Use(gin.Logger(), gin.CustomRecovery(middlewares.RecoveryErrors) /*, middlewares.TestMiddleware()*/)
}
