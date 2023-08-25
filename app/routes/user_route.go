package routes

import (
	"github.com/Firgisotya/go-rest-api/app/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoute(r *gin.Engine) {
	user := r.Group("/api/user/")
	{
		user.GET("", controllers.GetAllUser)
		user.POST("", controllers.CreateUser)
		user.GET(":id", controllers.ShowUser)
		user.PUT(":id", controllers.UpdateUser)
		user.DELETE(":id", controllers.DeleteUser)
	}
}