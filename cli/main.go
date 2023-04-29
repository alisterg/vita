package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	// "github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func main() {
	fmt.Println("Hello, World!")
}

func createSession() aws.Config {
	sess, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("<your AWS region>"),
	)
	if err != nil {
		// handle error
	}

	return sess
}
