package main

import (
	"log"
	"top-up-service/app"
	"top-up-service/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	app := SetupRouter()
	log.Fatal(app.Run(":8090"))
}

func SetupRouter() *gin.Engine {
	db := app.ConnectDatabase()
	app := gin.Default()

	app.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		AllowWildcard: true,
	}))

	routes.InitBalanceRoutes(db, app)

	return app
}
