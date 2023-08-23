package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {	
	AuthRoute(router)
	CategoryRoute(router)
	BooksRoute(router)
}
