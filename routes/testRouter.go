package routes

import (
	"communications_backend/controllers"
	"github.com/gin-gonic/gin"
)

func TestRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/index", controllers.TestController())
}
