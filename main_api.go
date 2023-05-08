//go:build api

package main

import (
	"github.com/aws/aws-lambda-go/lambda"

	"vita/transport/api"
)

func main() {
	lambda.Start(api.LambdaHandler)
}
