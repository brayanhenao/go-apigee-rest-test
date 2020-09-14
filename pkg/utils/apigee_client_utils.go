package utils

import (
	"context"
	"fmt"
	"golang.org/x/oauth2"
	"google.golang.org/api/apigee/v1"
	"google.golang.org/api/option"
	"os"
)

func GetApigeeService() (*apigee.Service, error) {
	ctx := context.Background()
	token := GetOauthToken()
	apigeeService, err := apigee.NewService(ctx, option.WithTokenSource(oauth2.StaticTokenSource(token)))
	if err != nil {
		return nil, err
	}

	apigeeService.BasePath = "https://api.enterprise.apigee.com/"

	return apigeeService, nil
}

func GetParent() string {
	return fmt.Sprintf("organizations/%s", os.Getenv("APIGEE_ORGANIZATION"))
}
