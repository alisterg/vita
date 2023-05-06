package app

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"vita/core"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

var tableName = "LifeDataTable"

func GetDbClient() (*dynamodb.Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err
	}

	return dynamodb.NewFromConfig(cfg), nil
}

func InsertEntry(client *dynamodb.Client, entry core.Entry) error {
	entryDataJson, err := json.Marshal(entry.Data)
	if err != nil {
		fmt.Printf("Couldn't serialise item data: %v:\n %v", err, entry.Data)
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

	fmt.Println("Item added")
	return nil
}

func Query(client *dynamodb.Client, partitionKey string, partitionValue string) (*dynamodb.QueryOutput, error) {
	input := &dynamodb.QueryInput{
		TableName:              aws.String(tableName),
		KeyConditionExpression: aws.String("#pk = :pkval"),
		ExpressionAttributeNames: map[string]string{
			"#pk": partitionKey,
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":pkval": &types.AttributeValueMemberS{Value: partitionValue},
		},
	}

	result, err := client.Query(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetAllEntriesForType(client *dynamodb.Client, entryType string) ([]*core.Entry, error) {
	var entries []*core.Entry

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
			return nil, fmt.Errorf("failed to get page: %v", err)
		}

		for _, item := range output.Items {
			entry, err := entryFromDynamoRecord(item)
			if err != nil {
				return nil, err
			}
			entries = append(entries, entry)
		}
	}

	return entries, nil
}

func entryFromDynamoRecord(item map[string]types.AttributeValue) (*core.Entry, error) {
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

	entry := core.Entry{
		Uuid:      uuid,
		EntryType: entryType,
		Data:      dataValJson,
		CreatedAt: createdAt,
	}

	return &entry, nil
}
