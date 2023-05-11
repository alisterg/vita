package lambda_api

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
	case http.MethodPut:
		return handlePutRequest(event)
	case http.MethodDelete:
		return handleDeleteRequest(event)
	default:
		return events.APIGatewayProxyResponse{}, fmt.Errorf("method not allowed")
	}
}

func handleGetRequest(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// /entry/all

	// /entry-type/all

	// /routine/all

	response := events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("Request: %v", event),
		StatusCode: 200,
	}
	return response, nil
}

func handlePostRequest(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// /entry/bulk
	// body: []{entryType: string, entryData: map[string]string}

	// /entry
	// body: {entryType: string, entryData: map[string]string}

	// /entry-type
	// body: {key: string, prompts: []string}

	// /routine
	// body: {key: string, entryTypes: []string}

	response := events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("Request: %v", event),
		StatusCode: 200,
	}
	return response, nil
}

func handlePutRequest(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// /entry
	// body: Entry

	// /entry-type
	// body: EntryType

	// /routine
	// body: Routine

	response := events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("Request: %v", event),
		StatusCode: 200,
	}
	return response, nil
}

func handleDeleteRequest(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// /entry/{uuid}

	// /entry-type/{key}

	// /routine/{key}

	response := events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("Request: %v", event),
		StatusCode: 200,
	}
	return response, nil
}
