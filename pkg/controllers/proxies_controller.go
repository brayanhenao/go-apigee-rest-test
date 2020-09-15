package controllers

import (
	"fmt"
	"github.com/brayanhenao/go-rest-test/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UploadFileTest(c *gin.Context) {
	// Source
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	filename := fmt.Sprintf("proxies/%s", file.Filename)
	if err := c.SaveUploadedFile(file, filename); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}

	c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully", file.Filename))
}

func CreateProxy(c *gin.Context) {
	file, _ := c.FormFile("file")
	proxyName := c.Query("name")
	filename := fmt.Sprintf("proxies/%s", file.Filename)
	if err := c.SaveUploadedFile(file, filename); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}

	apigeeService, err := utils.GetApigeeService()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	proxyRevision, _, err := apigeeService.Proxies.Import(proxyName, filename)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok": fmt.Sprintf("Proxy created! Revision up : %s", proxyRevision.Revision.String()),
	})
}

func GetProxy(c *gin.Context) {
	name := c.Query("name")

	apigeeService, err := utils.GetApigeeService()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	proxy, _, err := apigeeService.Proxies.Get(name)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"revisions":      fmt.Sprintf("%v", proxy.Revisions),
		"name":           proxy.Name,
		"metadata":       proxy.MetaData,
		"latestRevision": proxy.Revisions[len(proxy.Revisions)-1],
	})
}
