//go:build lambda_api

package main

import (
	"github.com/aws/aws-lambda-go/lambda"

	"vita/transport/lambda_api"
)

func main() {
	lambda.Start(lambda_api.LambdaHandler)
}
