package dynamo

import (
	"context"
	"encoding/json"
	"errors"

	"vita/core/entities"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type RoutinesDynamoRepo struct{}

var routinesTableName = "LifeRoutineTable"

func (d RoutinesDynamoRepo) CreateRoutine(routine entities.Routine) error {
	client, err := getDbClient()
	if err != nil {
		return err
	}

	entryTypesJson, err := json.Marshal(routine.EntryTypes)
	if err != nil {
		return err
	}

	insertObj := map[string]types.AttributeValue{
		"key":         &types.AttributeValueMemberS{Value: routine.Key},
		"entry_types": &types.AttributeValueMemberS{Value: string(entryTypesJson)},
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(routinesTableName),
		Item:      insertObj,
	}

	_, err2 := client.PutItem(context.TODO(), input)
	if err2 != nil {
		return err2
	}

	return nil
}

func (d RoutinesDynamoRepo) UpdateRoutine(routine entities.Routine) error {
	client, err := getDbClient()
	if err != nil {
		return err
	}

	entryTypesJson, err := json.Marshal(routine.EntryTypes)
	if err != nil {
		return err
	}

	updateObj := map[string]types.AttributeValue{
		":entry_types": &types.AttributeValueMemberS{Value: string(entryTypesJson)},
	}

	updateExpression := "SET #entry_types = :entry_types"

	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(routinesTableName),
		Key: map[string]types.AttributeValue{
			"key": &types.AttributeValueMemberS{Value: routine.Key},
		},
		UpdateExpression:          aws.String(updateExpression),
		ExpressionAttributeNames:  map[string]string{"#entry_types": "entry_types"},
		ExpressionAttributeValues: updateObj,
	}

	_, err2 := client.UpdateItem(context.TODO(), input)
	if err2 != nil {
		return err2
	}

	return nil
}

func (d RoutinesDynamoRepo) GetRoutine(key string) (entities.Routine, error) {
	client, err := getDbClient()
	if err != nil {
		return entities.Routine{}, err
	}

	input := &dynamodb.GetItemInput{
		TableName: aws.String(routinesTableName),
		Key: map[string]types.AttributeValue{
			"key": &types.AttributeValueMemberS{Value: key},
		},
	}

	result, err := client.GetItem(context.TODO(), input)
	if err != nil {
		return entities.Routine{}, err
	}

	if len(result.Item) == 0 {
		return entities.Routine{}, errors.New("routine not found")
	}

	routine, err := d.routineFromDynamoRecord(result.Item)
	if err != nil {
		return entities.Routine{}, err
	}

	return *routine, nil
}

func (d RoutinesDynamoRepo) DeleteRoutine(key string) error {
	return errors.New("not implemented")
}

func (d RoutinesDynamoRepo) GetAllRoutines() ([]entities.Routine, error) {
	client, err := getDbClient()
	if err != nil {
		return nil, err
	}

	var routines []entities.Routine

	input := &dynamodb.ScanInput{
		TableName: aws.String(routinesTableName),
	}

	pages := dynamodb.NewScanPaginator(client, input)

	for pages.HasMorePages() {
		output, err := pages.NextPage(context.TODO())
		if err != nil {
			return nil, err
		}

		for _, item := range output.Items {
			routine, err := d.routineFromDynamoRecord(item)
			if err != nil {
				return nil, err
			}
			routines = append(routines, *routine)
		}
	}

	return routines, nil
}

func (d RoutinesDynamoRepo) routineFromDynamoRecord(item map[string]types.AttributeValue) (*entities.Routine, error) {
	keyAttr, found := item["key"]
	if !found {
		return nil, errors.New("couldn't find key attribute")
	}
	keyValue, ok := keyAttr.(*types.AttributeValueMemberS)
	if !ok {
		return nil, errors.New("couldn't get key attribute")
	}
	entryTypeKey := keyValue.Value

	entryTypesAttr, found := item["entry_types"]
	if !found {
		return nil, errors.New("couldn't find entry_types attribute")
	}
	entryTypesValue, ok := entryTypesAttr.(*types.AttributeValueMemberS)
	if !ok {
		return nil, errors.New("couldn't get entry_types attribute")
	}

	var entryTypesValJson []string
	json.Unmarshal([]byte(entryTypesValue.Value), &entryTypesValJson)

	routine := entities.Routine{
		Key:        entryTypeKey,
		EntryTypes: entryTypesValJson,
	}

	return &routine, nil
}
