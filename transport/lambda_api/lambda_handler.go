package lambda_api

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"vita/core/repositories"
	"vita/core/use_cases"
	"vita/datasources/dynamo"

	"github.com/aws/aws-lambda-go/events"
)

var entryTypeRepo repositories.EntryTypeRepository
var entryRepo repositories.EntryRepository
var routineRepo repositories.RoutineRepository

var corsHeaders = map[string]string{
	"Access-Control-Allow-Headers": "Content-Type",
	"Access-Control-Allow-Origin":  "*",
	"Access-Control-Allow-Methods": "*",
}

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
	var err error

	switch event.Path {
	case "/entry/all":
		jsonResult, err = use_cases.GetAllEntriesJson(entryRepo)

	case "/entry-type/all":
		jsonResult, err = use_cases.GetAllEntryTypesJson(entryTypeRepo)

	case "/routine/all":
		jsonResult, err = use_cases.GetAllRoutinesJson(routineRepo)

	default:
		return events.APIGatewayProxyResponse{StatusCode: 404}, fmt.Errorf("not found")
	}

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	response := events.APIGatewayProxyResponse{
		Body:       jsonResult,
		Headers:    corsHeaders,
		StatusCode: 200,
	}
	return response, nil
}

func handlePostRequest(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var err error

	switch event.Path {
	// body: []EntryDto
	case "/entry/bulk":
		err = use_cases.CreateBulkEntriesFromJson(entryRepo, event.Body)

	// body: EntryDto
	case "/entry":
		err = use_cases.CreateEntryFromJson(entryRepo, event.Body)

	// body: EntryType
	case "/entry-type":
		err = use_cases.CreateEntryTypeFromJson(entryTypeRepo, event.Body)

	// body: Routine
	case "/routine":
		err = use_cases.CreateRoutineFromJson(routineRepo, event.Body)

	default:
		return events.APIGatewayProxyResponse{StatusCode: 404}, fmt.Errorf("not found")
	}

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	response := events.APIGatewayProxyResponse{
		Headers:    corsHeaders,
		StatusCode: 201,
	}
	return response, nil
}

func handlePutRequest(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var err error

	switch event.Path {
	case "/entry":
		// body: EntryDto
		err = use_cases.UpdateEntryFromJson(entryRepo, event.Body)

	case "/entry-type":
		// body: EntryType
		err = use_cases.UpdateEntryTypeFromJson(entryTypeRepo, event.Body)

	case "/routine":
		// body: Routine
		err = use_cases.UpdateRoutineFromJson(routineRepo, event.Body)

	default:
		return events.APIGatewayProxyResponse{StatusCode: 404}, fmt.Errorf("not found")
	}

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	response := events.APIGatewayProxyResponse{
		Headers:    corsHeaders,
		StatusCode: 204,
	}

	return response, nil
}

func handleDeleteRequest(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	if strings.HasPrefix(event.Path, "/entry/") { // /entry/{uuid}
		// uuid := strings.TrimPrefix(event.Path, "/entry/")
		// TODO: implement delete method. Dynamo requires composite key (including date),
		// so there isn't an easy way to do it

	} else if strings.HasPrefix(event.Path, "/entry-type/") { // /entry-type/{key}
		// key := strings.TrimPrefix(event.Path, "/entry-type/")
		// TODO: decide whether this cascade deletes entries

	} else if strings.HasPrefix(event.Path, "/routine/") { // /routine/{key}
		key := strings.TrimPrefix(event.Path, "/routine/")
		err := use_cases.DeleteRoutine(routineRepo, key)
		if err != nil {
			return events.APIGatewayProxyResponse{}, err
		}

	} else {
		return events.APIGatewayProxyResponse{StatusCode: 404}, fmt.Errorf("not found")
	}

	response := events.APIGatewayProxyResponse{
		Headers:    corsHeaders,
		StatusCode: 204,
	}
	return response, nil
}
