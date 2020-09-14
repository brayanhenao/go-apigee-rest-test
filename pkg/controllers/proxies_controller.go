package controllers

import (
	"fmt"
	"github.com/brayanhenao/go-rest-test/pkg/utils"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/apigee/v1"
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

	apiCall := apigeeService.Organizations.Apis.Create(utils.GetParent(), &apigee.GoogleApiHttpBody{
		ContentType: "multipart/form-data",
	}).Action("import").Name("test-go")

	proxyRevision, err := apiCall.Do()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok": fmt.Sprintf("Proxy created! Revision up : %s", proxyRevision.Revision),
	})
}
