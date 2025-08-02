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
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gorillamux"
)

// @title		Api
// @version		1.0
// @description	This is a sample server celler server.
// @termsOfService	http://swagger.io/terms/
func main() {

	router := server.Router()

	if strings.TrimSpace(os.Getenv("_LAMBDA_SERVER_PORT")) != "" {
		fmt.Println("Running in AWS Lambda environment")
		lambda.Start(func(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
			return gorillamux.NewV2(router).ProxyWithContext(ctx, request)
		})
	} else {
		fmt.Println("Running in local environment")
		if err := http.ListenAndServe(":3000", router); err != nil {
			fmt.Printf("Failed to start server: %v\n", err)
		}
	}

}
