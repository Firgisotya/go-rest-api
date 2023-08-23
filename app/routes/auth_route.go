package routes

import (
	"github.com/Firgisotya/go-rest-api/app/controllers"
	"github.com/Firgisotya/go-rest-api/app/middlewares"
	"github.com/gin-gonic/gin"
)

func AuthRoute(router *gin.Engine) {
	auth := router.Group("/api/auth")
	{
			auth.POST("/register", controllers.Register)
			auth.POST("/login", controllers.Login)
			auth.Use(middlewares.JWTMiddleware())
			auth.POST("/logout", controllers.Logout)
		
	}
}