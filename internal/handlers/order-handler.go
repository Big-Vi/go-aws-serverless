package handlers

import (
	"encoding/json"
	
	"github.com/aws/aws-lambda-go/events"
	"github.com/big-vi/go-aws-serverless/internal/models"
)

func GetOrders(req events.APIGatewayProxyRequest, DynamoDBTable string) (events.APIGatewayProxyResponse, error) {
	orders := models.GetOrders(DynamoDBTable)
	res, _ := json.Marshal(orders)
	return events.APIGatewayProxyResponse{Body: string(res), StatusCode: 200}, nil
}

func GetOrderById(req events.APIGatewayProxyRequest, DynamoDBTable string) (events.APIGatewayProxyResponse, error) {
	id := req.PathParameters["id"]
	order := models.GetOrderById(id, DynamoDBTable)

	res, _ := json.Marshal(order)
	return events.APIGatewayProxyResponse{Body: string(res), StatusCode: 200}, nil
}

func CreateOrder(req events.APIGatewayProxyRequest, DynamoDBTable string) (events.APIGatewayProxyResponse, error) {
	models.CreateOrder(DynamoDBTable)
	return events.APIGatewayProxyResponse{Body: "Order created", StatusCode: 200}, nil
}

func UpdateOrder(req events.APIGatewayProxyRequest, DynamoDBTable string) (events.APIGatewayProxyResponse, error) {
	id := req.QueryStringParameters["id"]
	models.UpdateOrder(id, DynamoDBTable)

	return events.APIGatewayProxyResponse{Body: "Order updated", StatusCode: 200}, nil
}

func DeleteOrder(req events.APIGatewayProxyRequest, DynamoDBTable string) (events.APIGatewayProxyResponse, error) {
	id := req.PathParameters["id"]
	models.DeleteOrder(id, DynamoDBTable)
	return events.APIGatewayProxyResponse{Body: "Order deleted", StatusCode: 200}, nil
}
