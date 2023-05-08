package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func LambdaHandler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch event.HTTPMethod {
	case http.MethodGet:
		return handleGetRequest(event)
	case http.MethodPost:
		return handlePostRequest(event)
	default:
		return events.APIGatewayProxyResponse{}, fmt.Errorf("method not allowed")
	}
}

func handleGetRequest(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	response := events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("Request: %v", event),
		StatusCode: 200,
	}
	return response, nil
}

func handlePostRequest(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	response := events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("Request: %v", event),
		StatusCode: 200,
	}
	return response, nil
}
