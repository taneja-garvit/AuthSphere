package routes

import (
	"login-signup-gin/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default() // Create Gin router with default middleware (logging, recovery)
	router.POST("/signup", controllers.Signup)
	router.POST("/login", controllers.Login)

	return router
}
