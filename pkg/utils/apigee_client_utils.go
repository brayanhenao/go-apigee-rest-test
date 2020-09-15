package utils

import (
	"github.com/zambien/go-apigee-edge"
	"log"
	"os"
)

func GetApigeeService() (*apigee.EdgeClient, error) {
	auth := GetAuth()
	opts := apigee.EdgeClientOptions{
		MgmtUrl: "https://api.enterprise.apigee.com/",
		Org:     os.Getenv("APIGEE_ORGANIZATION"),
		Auth:    &auth,
		Debug:   false,
	}

	client, err := apigee.NewEdgeClient(&opts)
	if err != nil {
		log.Printf("while initializing Edge client, error:\n%#v\n", err)
		return client, err
	}

	return client, nil
}
