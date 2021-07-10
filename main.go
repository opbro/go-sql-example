package main

import (
	"people/config"
	"people/controllers"
	"people/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Init Router
	db := config.Connect()

	controllers.CreatePersonTable(db)
	controllers.InitiateDB(db)

	router := gin.Default()

	routes.Routes(router)

	router.Run(":4747")
}
