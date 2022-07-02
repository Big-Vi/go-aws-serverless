package main

import(

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	response := events.APIGatewayProxyResponse{}
	switch req.HTTPMethod {
	case "GET":
		response = events.APIGatewayProxyResponse{Body: "Hellooo", StatusCode: 200}
	}
	return response, nil
}

func main() {
	lambda.Start(HandleRequest)
}