package handlers

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Register(handler *ReportsHandlerImpl) *gin.Engine {
	route := gin.Default()
	route.Use(func(c *gin.Context) {
		c.Next()
	})

	r := route.Group("/api")
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	tradesGroup := r.Group("/trades")
	tradesGroup.GET("/latest-report", handler.HandleLatestTrades)

	return route
}
