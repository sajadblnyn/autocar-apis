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
