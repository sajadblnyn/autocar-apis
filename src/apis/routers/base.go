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
