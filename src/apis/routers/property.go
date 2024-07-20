package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/sajadblnyn/autocar-apis/apis/handlers"
	"github.com/sajadblnyn/autocar-apis/config"
)

func Property(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewPropertyHandler(cfg)

	r.POST("/", h.CreateProperty)
	r.PUT("/:id", h.UpdateProperty)
	r.GET("/:id", h.GetProperty)
	r.DELETE("/:id", h.DeleteProperty)
	r.POST("/filter", h.GetPropertiesByFilter)

}

func PropertyCategory(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewPropertyCategoryHandler(cfg)

	r.POST("/", h.CreatePropertyCategory)
	r.PUT("/:id", h.UpdatePropertyCategory)
	r.GET("/:id", h.GetPropertyCategory)
	r.DELETE("/:id", h.DeletePropertyCategory)
	r.POST("/filter", h.GetPropertyCategoriesByFilter)

}
