package main

import (
	"context"
	/*"log"
	"os"

	"github.com/SimonMora/bikesams_be/aws_go"
	"github.com/SimonMora/bikesams_be/database"
	"github.com/SimonMora/bikesams_be/models"*/

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(ExecuteLambda)
}

func ExecuteLambda(context context.Context, request events.APIGatewayV2HTTPRequest) (*events.APIGatewayProxyResponse, error) {

}
