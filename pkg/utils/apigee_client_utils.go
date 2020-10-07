package utils

import (
	"github.com/brayanhenao/go-apigee-edge/apigee"
	"log"
	"os"
)

func GetApigeeService() (*apigee.ApigeeClient, error) {
	auth := GetAuth()
	opts := apigee.ApigeeClientOptions{
		MgmtUrl: "https://api.enterprise.apigee.com/",
		Org:     os.Getenv("APIGEE_ORGANIZATION"),
		Auth:    &auth,
		Debug:   false,
	}

	client, err := apigee.NewApigeeClient(&opts)
	if err != nil {
		log.Printf("while initializing Edge client, error:\n%#v\n", err)
		return client, err
	}

	return client, nil
}
