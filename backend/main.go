package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApplicationRequest struct {
	Name string `json:"name"`

	Image string `json:"image"`

	Replicas int32 `json:"replicas"`

	CPU string `json:"cpu"`

	Memory string `json:"memory"`
}

func main() {

	router := gin.Default()

	router.POST("/applications", func(ctx *gin.Context) {

		var app ApplicationRequest

		if err := ctx.ShouldBindJSON(&app); err != nil {

			ctx.JSON(
				http.StatusBadRequest,
				gin.H{
					"error": err.Error(),
				},
			)

			return

		}

		// Temporary: print received application

		println(
			"Creating application:",
			app.Name,
		)

		ctx.JSON(
			http.StatusCreated,
			app,
		)

	})

	router.Run(":8080")

}
