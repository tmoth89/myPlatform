package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/tmoth89/myPlatform/backend/models"
	"github.com/tmoth89/myPlatform/backend/services"
)

func CreateApplication(c *gin.Context) {

	var app models.ApplicationRequest

	if err := c.ShouldBindJSON(&app); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	err := services.CreateApplication(app)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Application created successfully",
	})
}