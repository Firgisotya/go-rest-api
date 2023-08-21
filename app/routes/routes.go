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

	todo := router.Group("/todo")
	{
		todo.Use(middlewares.JWTMiddleware())
		todo.POST("/", controllers.CreateTodo)
		todo.GET("/", controllers.GetTodos)
		todo.PUT("/:id", controllers.UpdateTodo)
		todo.DELETE("/:id", controllers.DeleteTodo)
		
	}
	}

	
	return r
}
