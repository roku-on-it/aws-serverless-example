package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type headers map[string]string

func handler() (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body: "Hello World! I am written in Go.",
		Headers: headers{
			"Content-Type": "text/plain",
		},
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
