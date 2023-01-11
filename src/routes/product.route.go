package routes

import (
	"github.com/Victor1890/gin-crud-api-rest/src/controllers"
	"github.com/gin-gonic/gin"
)

func ProductGroup(router *gin.RouterGroup) {

	router.GET("/products", controllers.GetAllProductHandler)
	router.GET("/products/", controllers.GetAllProductHandler)
	router.GET("/products/:id", controllers.GetOneProductHandler)
	router.POST("/products", controllers.CreateProductHandler)
	router.POST("/products/", controllers.CreateProductHandler)
	router.PATCH("/products/:id", controllers.UpdateOneProductHandler)
	router.DELETE("/products/:id", controllers.DeleteOneProductHandler)

}
