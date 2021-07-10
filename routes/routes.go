package routes

import (
	"net/http"
	"people/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.GET("/people", controllers.GetAllPeople)
	router.POST("/person", controllers.CreatePerson)
}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome to API",
	})
	return
}
