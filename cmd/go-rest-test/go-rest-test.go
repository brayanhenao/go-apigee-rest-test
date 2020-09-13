package main

import (
	"github.com/brayanhenao/go-rest-test/pkg/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.RegisterRoutes(r)
	r.Run()
}
