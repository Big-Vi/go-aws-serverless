package main

import (
	"github.com/big-vi/go-aws-serverless/internal/handlers"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var DynamoDBTable string = "CRUDDynamoDBTable"

func HandleRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch req.HTTPMethod {
	case "GET":
		if req.PathParameters["id"] != "" {
			return handlers.GetOrderById(req, DynamoDBTable)
		}
		return handlers.GetOrders(req, DynamoDBTable)
	case "PUT":
		if req.QueryStringParameters["id"] != "" {
			return handlers.UpdateOrder(req, DynamoDBTable)
		}
		return handlers.CreateOrder(req, DynamoDBTable)
	case "DELETE":
		return handlers.DeleteOrder(req, DynamoDBTable)
	default:
		return events.APIGatewayProxyResponse{Body: "HTTP request is not valid.", StatusCode: 404}, nil
	}
}

func main() {
	lambda.Start(HandleRequest)
}
