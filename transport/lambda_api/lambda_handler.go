package lambda_api

import (
	"context"
	"fmt"
	"net/http"

	"vita/core/repositories"
	"vita/core/use_cases"
	"vita/datasources/dynamo"

	"github.com/aws/aws-lambda-go/events"
)

var entryTypeRepo repositories.EntryTypeRepository
var entryRepo repositories.EntryRepository
var routineRepo repositories.RoutineRepository

func LambdaHandler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// CHANGE THESE IF YOU WISH TO CHANGE THE DATA SOURCE
	entryTypeRepo = dynamo.EntryTypesDynamoRepo{}
	routineRepo = dynamo.RoutinesDynamoRepo{}
	entryRepo = dynamo.EntriesDynamoRepo{}

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
	var jsonResult string

	switch event.Path {
	case "/entry/all":
		entries, err := use_cases.GetAllEntriesJson(entryRepo)
		if err != nil {
			return events.APIGatewayProxyResponse{}, err
		}
		jsonResult = entries

	case "/entry-type/all":
	case "/routine/all":
	default:
		return events.APIGatewayProxyResponse{StatusCode: 404}, fmt.Errorf("not found")
	}

	response := events.APIGatewayProxyResponse{
		Body:       jsonResult,
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
