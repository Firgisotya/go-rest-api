package routes

import (
	"github.com/Firgisotya/go-rest-api/app/controllers"
	"github.com/Firgisotya/go-rest-api/app/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	router := r.Group("/api")
	{
		auth := router.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
	}

	category := router.Group("/category")
	{
		category.Use(middlewares.JWTMiddleware())
		category.GET("/", controllers.GetAllCategory)
	}
	}

	
	return r
}
