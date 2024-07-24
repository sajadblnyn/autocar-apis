package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/sajadblnyn/autocar-apis/apis/handlers"
	"github.com/sajadblnyn/autocar-apis/config"
)

func CarType(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewCarTypeHandler(cfg)

	r.POST("/", h.CreateCarType)
	r.PUT("/:id", h.UpdateCarType)
	r.GET("/:id", h.GetCarType)
	r.DELETE("/:id", h.DeleteCarType)
	r.POST("/filter", h.GetCarTypesByFilter)

}

func Gearbox(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewGearboxHandler(cfg)

	r.POST("/", h.CreateGearbox)
	r.PUT("/:id", h.UpdateGearbox)
	r.GET("/:id", h.GetGearbox)
	r.DELETE("/:id", h.DeleteGearbox)
	r.POST("/filter", h.GetGearboxesByFilter)

}

func Company(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewCompanyHandler(cfg)

	r.POST("/", h.CreateCompany)
	r.PUT("/:id", h.UpdateCompany)
	r.GET("/:id", h.GetCompany)
	r.DELETE("/:id", h.DeleteCompany)
	r.POST("/filter", h.GetCompaniesByFilter)

}

func CarModel(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewCarModelHandler(cfg)

	r.POST("/", h.CreateCarModel)
	r.PUT("/:id", h.UpdateCarModel)
	r.GET("/:id", h.GetCarModel)
	r.DELETE("/:id", h.DeleteCarModel)
	r.POST("/filter", h.GetCarModelsByFilter)

}

func CarModelColor(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewCarModelColorHandler(cfg)

	r.POST("/", h.CreateCarModelColor)
	r.PUT("/:id", h.UpdateCarModelColor)
	r.GET("/:id", h.GetCarModelColor)
	r.DELETE("/:id", h.DeleteCarModelColor)
	r.POST("/filter", h.GetCarModelColorsByFilter)

}

func CarModelPersianYear(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewCarModelPersianYearHandler(cfg)

	r.POST("/", h.CreateCarModelPersianYear)
	r.PUT("/:id", h.UpdateCarModelPersianYear)
	r.GET("/:id", h.GetCarModelPersianYear)
	r.DELETE("/:id", h.DeleteCarModelPersianYear)
	r.POST("/filter", h.GetCarModelPersianYearsByFilter)

}

func CarModelPriceHistory(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewCarModelPriceHistoryHandler(cfg)

	r.POST("/", h.CreateCarModelPriceHistory)
	r.PUT("/:id", h.UpdateCarModelPriceHistory)
	r.GET("/:id", h.GetCarModelPriceHistory)
	r.DELETE("/:id", h.DeleteCarModelPriceHistory)
	r.POST("/filter", h.GetCarModelPriceHistoriesByFilter)

}

func CarModelImage(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewCarModelImageHandler(cfg)

	r.POST("/", h.CreateCarModelImage)
	r.PUT("/:id", h.UpdateCarModelImage)
	r.GET("/:id", h.GetCarModelImage)
	r.DELETE("/:id", h.DeleteCarModelImage)
	r.POST("/filter", h.GetCarModelImagesByFilter)

}

func CarModelProperty(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewCarModelPropertyHandler(cfg)

	r.POST("/", h.CreateCarModelProperty)
	r.PUT("/:id", h.UpdateCarModelProperty)
	r.GET("/:id", h.GetCarModelProperty)
	r.DELETE("/:id", h.DeleteCarModelProperty)
	r.POST("/filter", h.GetCarModelPropertiesByFilter)

}
