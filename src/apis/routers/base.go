package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/sajadblnyn/autocar-apis/apis/handlers"
	"github.com/sajadblnyn/autocar-apis/config"
)

func Country(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewCountryHandler(cfg)

	r.POST("/", h.CreateCountry)
	r.PUT("/:id", h.UpdateCountry)
	r.GET("/:id", h.GetCountry)
	r.DELETE("/:id", h.DeleteCountry)
	r.POST("/filter", h.GetCountriesByFilter)

}

func City(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewCityHandler(cfg)

	r.POST("/", h.CreateCity)
	r.PUT("/:id", h.UpdateCity)
	r.GET("/:id", h.GetCity)
	r.DELETE("/:id", h.DeleteCity)
	r.POST("/filter", h.GetCitiesByFilter)

}

func File(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewFileHandler(cfg)

	r.POST("/", h.CreateFile)
	r.PUT("/:id", h.UpdateFile)
	r.GET("/:id", h.GetFile)
	r.DELETE("/:id", h.DeleteFile)
	r.POST("/filter", h.GetFilesByFilter)

}

func Color(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewColorHandler(cfg)

	r.POST("/", h.CreateColor)
	r.PUT("/:id", h.UpdateColor)
	r.GET("/:id", h.GetColor)
	r.DELETE("/:id", h.DeleteColor)
	r.POST("/filter", h.GetColorsByFilter)

}

func PersianYear(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewPersianYearHandler(cfg)

	r.POST("/", h.CreatePersianYear)
	r.PUT("/:id", h.UpdatePersianYear)
	r.GET("/:id", h.GetPersianYear)
	r.DELETE("/:id", h.DeletePersianYear)
	r.POST("/filter", h.GetPersianYearsByFilter)

}
