package config

import (
	"github.com/aws/aws-sdk-go-v2/config"
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func ConnectDynamoDB() *dynamodb.Client  {
	cfg, err := config.LoadDefaultConfig(context.TODO(), func(opts *config.LoadOptions) error {
		opts.Region = "ap-southeast-2"
		return nil
	})
	if err != nil {
		panic(err)
	}

	dynamodbClient := dynamodb.NewFromConfig(cfg)
	return dynamodbClient
}
