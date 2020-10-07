package utils

import (
	"github.com/brayanhenao/go-apigee-edge/apigee"
	"os"
)

func GetAuth() apigee.AdminAuth {
	return apigee.AdminAuth{
		Username: os.Getenv("APIGEE_USERNAME"),
		Password: os.Getenv("APIGEE_password"),
	}
}
