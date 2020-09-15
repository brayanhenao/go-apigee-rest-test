package utils

import (
	"github.com/zambien/go-apigee-edge"
	"os"
)

func GetAuth() apigee.EdgeAuth {
	return apigee.EdgeAuth{
		Username: os.Getenv("APIGEE_USERNAME"),
		Password: os.Getenv("APIGEE_password"),
	}
}
