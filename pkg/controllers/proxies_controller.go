package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func CreateProxy(c *gin.Context) {
	file, _ := c.FormFile("file")

	c.JSON(200, gin.H{
		"message": fmt.Sprintf("File %s uploaded", file.Filename),
	})
}
