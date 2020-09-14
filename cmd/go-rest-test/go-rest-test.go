package main

import (
	"encoding/json"
	"fmt"
	"github.com/brayanhenao/go-rest-test/pkg/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"google.golang.org/api/apigee/v1"
	"log"
	"os"
	"time"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	res, _ := os.Open("test.json")

	object := apigee.GoogleCloudApigeeV1ApiProxy{}
	err = json.NewDecoder(res).Decode(&object)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(object)

	defer res.Close()

	gin.ForceConsoleColor()
	server := gin.New()
	server.Use(gin.Recovery())
	server.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	routes.RegisterRoutes(server)
	_ = server.Run()
}
