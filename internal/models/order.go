package models

import (
	"context"
	"fmt"

	"github.com/big-vi/go-aws-serverless/config"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
)

var dynamodbClient *dynamodb.Client

func init() {
	dynamodbClient = config.ConnectDynamoDB()
}

type Order struct {
	ID    string
	Title string
}

func GetOrders(DynamoDBTable string) []Order {
	result, err := dynamodbClient.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName: aws.String(DynamoDBTable),
	})
	if err != nil {
		panic(err)
	}

	var orders []Order

	//Using Go SDK helper function to parse dynamodb attributes to struct
	err = attributevalue.UnmarshalListOfMaps(result.Items, &orders)
	if err != nil {
		err = fmt.Errorf("failed to marshal Record, %w", err)
		fmt.Println("An error happened:", err)
	}

	return orders
}

func GetOrderById(id string, DynamoDBTable string) Order {
	result, err := dynamodbClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(DynamoDBTable),
		Key: map[string]types.AttributeValue{
			"ID": &types.AttributeValueMemberS{Value: id},
		},
	})

	if err != nil {
		panic(err)
	}

	var order Order

	err = attributevalue.UnmarshalMap(result.Item, &order)
	if err != nil {
		err = fmt.Errorf("failed to marshal Record, %w", err)
		fmt.Println("An error happened:", err)
	}

	return order
}

func DeleteOrder(id string, DynamoDBTable string) *dynamodb.DeleteItemOutput {
	result, err := dynamodbClient.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
		TableName: aws.String(DynamoDBTable),
		Key: map[string]types.AttributeValue{
			"ID": &types.AttributeValueMemberS{Value: id},
		},
	})
	if err != nil {
		panic(err)
	}

	return result
}

func CreateOrder(DynamoDBTable string) *dynamodb.PutItemOutput {
	id := uuid.New()
	result, err := dynamodbClient.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(DynamoDBTable),
		Item: map[string]types.AttributeValue{
			"ID":    &types.AttributeValueMemberS{Value: id.String()},
			"Title": &types.AttributeValueMemberS{Value: "John Doe"},
		},
	})

	if err != nil {
		panic(err)
	}

	return result
}

func UpdateOrder(id string, DynamoDBTable string) *dynamodb.PutItemOutput {
	result, err := dynamodbClient.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(DynamoDBTable),
		Item: map[string]types.AttributeValue{
			"ID": &types.AttributeValueMemberS{Value: id},
			"Title": &types.AttributeValueMemberS{Value: "Peter"},
		},
	})

	if err != nil {
		panic(err)
	}

	return result
}
