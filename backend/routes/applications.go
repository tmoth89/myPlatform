package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/tmoth89/myPlatform/backend/handlers"
)

func RegisterRoutes(router *gin.Engine) {

	router.GET("/", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"message": "Platform API Running",
		})

	})

	router.POST("/applications", handlers.CreateApplication)
}