package main

import (
	"github.com/Victor1890/gin-crud-api-rest/src/config"
	"github.com/Victor1890/gin-crud-api-rest/src/database"
	"github.com/Victor1890/gin-crud-api-rest/src/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	var config = config.Config()

	database.Database()
	var router = gin.New()

	gin.SetMode(config.ReleaseMode)

	router.SetTrustedProxies([]string{"::1"})
	router.Static("/assets", "./assets")

	var apiRouter = router.Group("/api")
	routes.ProductGroup(apiRouter)

	router.Run(":3000")
}
