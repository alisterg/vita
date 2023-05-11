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

type EntryTypesDynamoRepo struct{}

var entryTypeTableName = "LifeEntryTypeTable"

func (d EntryTypesDynamoRepo) CreateEntryType(entryType entities.EntryType) error {
	client, err := getDbClient()
	if err != nil {
		return err
	}

	promptsJson, err := json.Marshal(entryType.Prompts)
	if err != nil {
		return err
	}

	insertObj := map[string]types.AttributeValue{
		"key":     &types.AttributeValueMemberS{Value: entryType.Key},
		"prompts": &types.AttributeValueMemberS{Value: string(promptsJson)},
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(entryTypeTableName),
		Item:      insertObj,
	}

	_, err2 := client.PutItem(context.TODO(), input)
	if err2 != nil {
		return err2
	}

	return nil
}

func (d EntryTypesDynamoRepo) UpdateEntryType(entryType entities.EntryType) error {
	client, err := getDbClient()
	if err != nil {
		return err
	}

	promptsJson, err := json.Marshal(entryType.Prompts)
	if err != nil {
		return err
	}

	updateObj := map[string]types.AttributeValue{
		":prompts": &types.AttributeValueMemberS{Value: string(promptsJson)},
	}

	updateExpression := "SET #prompts = :prompts"

	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(entryTypeTableName),
		Key: map[string]types.AttributeValue{
			"key": &types.AttributeValueMemberS{Value: entryType.Key},
		},
		UpdateExpression:          aws.String(updateExpression),
		ExpressionAttributeNames:  map[string]string{"#prompts": "prompts"},
		ExpressionAttributeValues: updateObj,
	}

	_, err2 := client.UpdateItem(context.TODO(), input)
	if err2 != nil {
		return err2
	}

	return nil
}

func (d EntryTypesDynamoRepo) GetEntryType(key string) (entities.EntryType, error) {
	client, err := getDbClient()
	if err != nil {
		return entities.EntryType{}, err
	}

	input := &dynamodb.GetItemInput{
		TableName: aws.String(entryTypeTableName),
		Key: map[string]types.AttributeValue{
			"key": &types.AttributeValueMemberS{Value: key},
		},
	}

	result, err := client.GetItem(context.TODO(), input)
	if err != nil {
		return entities.EntryType{}, err
	}

	if len(result.Item) == 0 {
		return entities.EntryType{}, errors.New("entry type not found")
	}

	entryType, err := d.entryTypeFromDynamoRecord(result.Item)
	if err != nil {
		return entities.EntryType{}, err
	}

	return *entryType, nil
}

func (d EntryTypesDynamoRepo) DeleteEntryType(key string) error {
	return errors.New("not implemented")
}

func (d EntryTypesDynamoRepo) GetAllEntryTypes() ([]entities.EntryType, error) {
	client, err := getDbClient()
	if err != nil {
		return nil, err
	}

	var entryTypes []entities.EntryType

	input := &dynamodb.ScanInput{
		TableName: aws.String(entryTypeTableName),
	}

	pages := dynamodb.NewScanPaginator(client, input)

	for pages.HasMorePages() {
		output, err := pages.NextPage(context.TODO())
		if err != nil {
			return nil, err
		}

		for _, item := range output.Items {
			entryType, err := d.entryTypeFromDynamoRecord(item)
			if err != nil {
				return nil, err
			}
			entryTypes = append(entryTypes, *entryType)
		}
	}

	return entryTypes, nil
}

func (d EntryTypesDynamoRepo) entryTypeFromDynamoRecord(item map[string]types.AttributeValue) (*entities.EntryType, error) {
	keyAttr, found := item["key"]
	if !found {
		return nil, errors.New("couldn't find key attribute")
	}
	keyValue, ok := keyAttr.(*types.AttributeValueMemberS)
	if !ok {
		return nil, errors.New("couldn't get key attribute")
	}
	entryTypeKey := keyValue.Value

	promptsAttr, found := item["prompts"]
	if !found {
		return nil, errors.New("couldn't find prompts attribute")
	}
	promptsValue, ok := promptsAttr.(*types.AttributeValueMemberS)
	if !ok {
		return nil, errors.New("couldn't get prompts attribute")
	}

	var promptsValJson []string
	json.Unmarshal([]byte(promptsValue.Value), &promptsValJson)

	entryType := entities.EntryType{
		Key:     entryTypeKey,
		Prompts: promptsValJson,
	}

	return &entryType, nil
}
