package routes

import (
	"github.com/Victor1890/gin-crud-api-rest/src/controllers"
	"github.com/gin-gonic/gin"
)

func ProductGroup(router *gin.RouterGroup) {

	router.GET("", controllers.GetAllProductHandler)
	router.GET("/", controllers.GetAllProductHandler)
	router.GET("/:id", controllers.GetOneProductHandler)
	router.POST("", controllers.CreateProductHandler)
	router.PATCH("/:id", controllers.UpdateOneProductHandler)
	router.DELETE("/:id", controllers.DeleteOneProductHandler)

}
