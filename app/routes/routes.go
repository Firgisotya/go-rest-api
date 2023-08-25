package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {	
	AuthRoute(router)
	UserRoute(router)
	CategoryRoute(router)
	BooksRoute(router)
}
