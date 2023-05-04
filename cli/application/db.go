package application

import (
	"context"
	"encoding/json"
	"fmt"

	"vita/core"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

var tableName = "lifedata"

func GetDbClient() (*dynamodb.Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err
	}

	return dynamodb.NewFromConfig(cfg), nil
}

func InsertEntry(client *dynamodb.Client, item core.Entry) error {
	itemData, err := json.Marshal(item.ItemData)
	if err != nil {
		fmt.Printf("Couldn't serialise item data: %v:\n %v", err, item.ItemData)
		return err
	}

	insertObj := map[string]types.AttributeValue{
		"type": &types.AttributeValueMemberS{Value: item.ItemType},
		"date": &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", item.ItemDate)},
		"data": &types.AttributeValueMemberS{Value: string(itemData)},
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

func GenericQuery(client *dynamodb.Client, table string, partitionKey string, partitionValue string) (*dynamodb.QueryOutput, error) {
	input := &dynamodb.QueryInput{
		TableName:              aws.String(table),
		KeyConditionExpression: aws.String("#pk = :pkval"),
		ExpressionAttributeNames: map[string]string{
			"#pk": partitionKey,
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":pkval": &types.AttributeValueMemberS{Value: partitionValue},
		},
	}

	result, err := client.Query(context.Background(), input)
	if err != nil {
		return nil, err
	}

	return result, nil
}
