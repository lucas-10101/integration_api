// swagger:meta
//
// Package main My API.
//
// This is a sample server.
// Schemes: http, https
// Host: localhost:8080
// BasePath: /api/v1
// Version: 1.0.0
// Contact: Your Name <you@example.com>
//
//go:generate swagger generate spec -o swagger.json
package main

import (
	"api/server"
	"api/services"
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gorillamux"
)

func main() {

	var i18n = services.I18nService{}
	i18n.Init()

	var languagePack = i18n.GetLanguagePack(services.LANG_EN)

	fmt.Println(languagePack.T("hello"))

	return

	router := server.Router()

	runMethod := os.Getenv("RUN_METHOD")

	switch runMethod {
	case "local":
		fmt.Println("Running in local environment")
		if err := http.ListenAndServe(":3000", router); err != nil {
			fmt.Printf("Failed to start server: %v\n", err)
		}
	case "lambda":
		fmt.Println("Running in AWS Lambda environment")

		lambda.Start(func(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
			return gorillamux.NewV2(router).ProxyWithContext(ctx, request)
		})
	default:
		fmt.Println("No valid run method specified, aborting execution")
		os.Exit(1)
	}
}
