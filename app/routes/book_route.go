package routes

import (
	"github.com/Firgisotya/go-rest-api/app/controllers"
	// "github.com/Firgisotya/go-rest-api/app/middlewares"
	"github.com/gin-gonic/gin"
)

func BooksRoute(router *gin.Engine) {
	book := router.Group("/api/book/")
	{
		// book.Use(middlewares.JWTMiddleware())
		book.GET("", controllers.GetAllBook)
		book.POST("", controllers.CreateBook)
		book.GET(":id", controllers.ShowBook)
		book.PUT(":id", controllers.UpdateBook)
		book.DELETE(":id", controllers.DeleteBook)
	}
}
