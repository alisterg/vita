package dynamo

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"vita/core/entities"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type EntriesDynamoRepo struct{}

var tableName = "LifeDataTable"

func (d EntriesDynamoRepo) getDbClient() (*dynamodb.Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err
	}

	return dynamodb.NewFromConfig(cfg), nil
}

func (d EntriesDynamoRepo) CreateEntry(entry entities.Entry) error {
	client, err := d.getDbClient()
	if err != nil {
		return err
	}

	entryDataJson, err := json.Marshal(entry.Data)
	if err != nil {
		return err
	}

	insertObj := map[string]types.AttributeValue{
		"uuid": &types.AttributeValueMemberS{Value: entry.Uuid},
		"type": &types.AttributeValueMemberS{Value: entry.EntryType},
		"date": &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", entry.CreatedAt)},
		"data": &types.AttributeValueMemberS{Value: string(entryDataJson)},
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      insertObj,
	}

	_, err2 := client.PutItem(context.TODO(), input)
	if err2 != nil {
		return err2
	}

	return nil
}

func (d EntriesDynamoRepo) UpdateEntry(entry entities.Entry) error {
	client, err := d.getDbClient()
	if err != nil {
		return err
	}

	entryDataJson, err := json.Marshal(entry.Data)
	if err != nil {
		return err
	}

	updateObj := map[string]types.AttributeValue{
		":data": &types.AttributeValueMemberS{Value: string(entryDataJson)},
	}

	updateExpression := "SET #data = :data"

	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(tableName),
		Key: map[string]types.AttributeValue{
			"uuid": &types.AttributeValueMemberS{Value: entry.Uuid},
			"date": &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", entry.CreatedAt)},
		},
		UpdateExpression:          aws.String(updateExpression),
		ExpressionAttributeNames:  map[string]string{"#data": "data"},
		ExpressionAttributeValues: updateObj,
	}

	_, err2 := client.UpdateItem(context.TODO(), input)
	if err2 != nil {
		return err2
	}

	return nil
}

func (d EntriesDynamoRepo) BulkCreateEntries(entries []entities.Entry) error {
	client, err := d.getDbClient()
	if err != nil {
		return err
	}

	writeRequests := make([]types.WriteRequest, len(entries))
	for i, entry := range entries {
		entryDataJson, err := json.Marshal(entry.Data)
		if err != nil {
			return err
		}

		insertObj := map[string]types.AttributeValue{
			"uuid": &types.AttributeValueMemberS{Value: entry.Uuid},
			"type": &types.AttributeValueMemberS{Value: entry.EntryType},
			"date": &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", entry.CreatedAt)},
			"data": &types.AttributeValueMemberS{Value: string(entryDataJson)},
		}

		writeRequests[i] = types.WriteRequest{
			PutRequest: &types.PutRequest{
				Item: insertObj,
			},
		}
	}

	batchWriteInput := &dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]types.WriteRequest{
			tableName: writeRequests,
		},
	}

	_, err2 := client.BatchWriteItem(context.TODO(), batchWriteInput)
	if err2 != nil {
		return err2
	}

	return nil
}

func (d EntriesDynamoRepo) GetAllEntriesForType(entryType string) ([]entities.Entry, error) {
	client, err := d.getDbClient()
	if err != nil {
		return nil, err
	}

	var entries []entities.Entry

	input := &dynamodb.ScanInput{
		TableName:        aws.String(tableName),
		FilterExpression: aws.String("#type = :type"),
		ExpressionAttributeNames: map[string]string{
			"#type": "type",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":type": &types.AttributeValueMemberS{Value: entryType},
		},
	}

	pages := dynamodb.NewScanPaginator(client, input)

	for pages.HasMorePages() {
		output, err := pages.NextPage(context.TODO())
		if err != nil {
			return nil, err
		}

		for _, item := range output.Items {
			entry, err := d.entryFromDynamoRecord(item)
			if err != nil {
				return nil, err
			}
			entries = append(entries, *entry)
		}
	}

	return entries, nil
}

func (d EntriesDynamoRepo) entryFromDynamoRecord(item map[string]types.AttributeValue) (*entities.Entry, error) {
	uuidAttr, found := item["uuid"]
	if !found {
		return nil, errors.New("couldn't find uuid attribute")
	}
	uuidValue, ok := uuidAttr.(*types.AttributeValueMemberS)
	if !ok {
		return nil, errors.New("couldn't get uuid attribute")
	}
	uuid := uuidValue.Value

	typeAttr, found := item["type"]
	if !found {
		return nil, errors.New("couldn't find type attribute")
	}
	typeValue, ok := typeAttr.(*types.AttributeValueMemberS)
	if !ok {
		return nil, errors.New("couldn't get type attribute")
	}
	entryType := typeValue.Value

	dataAttr, found := item["data"]
	if !found {
		return nil, errors.New("couldn't find data attribute")
	}
	dataValue, ok := dataAttr.(*types.AttributeValueMemberS)
	if !ok {
		return nil, errors.New("couldn't get data attribute")
	}

	var dataValJson map[string]string
	json.Unmarshal([]byte(dataValue.Value), &dataValJson)

	createdAtAttr, found := item["date"]
	if !found {
		return nil, errors.New("couldn't find date attribute")
	}
	createdAtValue, ok := createdAtAttr.(*types.AttributeValueMemberN)
	if !ok {
		return nil, errors.New("couldn't get date attribute")
	}
	createdAt, err := strconv.ParseInt(createdAtValue.Value, 10, 64)
	if err != nil {
		return nil, errors.New("couldn't parse date attribute")
	}

	entry := entities.Entry{
		Uuid:      uuid,
		EntryType: entryType,
		Data:      dataValJson,
		CreatedAt: createdAt,
	}

	return &entry, nil
}
