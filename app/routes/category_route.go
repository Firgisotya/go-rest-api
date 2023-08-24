package routes

import (
	"github.com/Firgisotya/go-rest-api/app/controllers"
	// "github.com/Firgisotya/go-rest-api/app/middlewares"
	"github.com/gin-gonic/gin"
)

func CategoryRoute(router *gin.Engine) {
	category := router.Group("/api/category/")
	{
		// category.Use(middlewares.JWTMiddleware())
		category.GET("", controllers.GetAllCategory)
		category.POST("", controllers.CreateCategory)
		category.PUT(":id", controllers.UpdateCategory)
		category.DELETE(":id", controllers.DeleteCategory)
	}
}
