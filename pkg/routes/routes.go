package routes

import (
	"github.com/brayanhenao/go-rest-test/pkg/controllers"
	"github.com/gin-gonic/gin"
)

var RegisterRoutes = func(server *gin.Engine) {
	v1 := server.Group("/api/v1")
	{
		v1.GET("/health", controllers.Health)
		v1.POST("/upload", controllers.CreateProxy)
	}
}
