package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"stori/pkg/domain"
)

func main() {
	lambda.Start(handler)
}

func handler(evt events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	err := domain.SendSummary(evt.QueryStringParameters["address"])

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       err.Error(),
		}, nil
	} else {
		response := events.APIGatewayProxyResponse{
			StatusCode: 200,
			Body:       "Summary email sent successfully.",
		}
		return response, nil
	}

}
