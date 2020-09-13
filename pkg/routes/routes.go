package routes

import (
	"github.com/brayanhenao/go-rest-test/pkg/controllers"
	"github.com/gin-gonic/gin"
)

var RegisterRoutes = func(server *gin.Engine) {
	server.GET("/health", controllers.Health)
}
