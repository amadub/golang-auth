package main

import (
	"amadub/golang-auth/controllers"
	"amadub/golang-auth/database"
	"amadub/golang-auth/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Database
	database.Connect("root:password@tcp(localhost:3306)/golang-auth?parseTime=true")
	database.Migrate()
	// Initialize Router
	router := initRouter()
	router.Run(":8080")
}
func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/token", controllers.GenerateToken)
		api.POST("/user/register", controllers.RegisterUser)
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Ping)
		}
	}
	return router
}
