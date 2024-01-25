package main

import (
	"context"
	"os"
	"strings"

	"github.com/SimonMora/bikesams_be/aws_go"
	"github.com/SimonMora/bikesams_be/handlers"
	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
)

const ENV_SECRET_NAME = "SecretName"
const ENV_URL_PREFIX = "UrlPrefix"

func main() {
	lambda.Start(ExecuteLambda)
}

func ExecuteLambda(context context.Context, request events.APIGatewayV2HTTPRequest) (*events.APIGatewayProxyResponse, error) {
	//Initialize the aws environment
	aws_go.InitAws()

	isParam, paramName := ValidateEnvironmentVariables()
	if !isParam && paramName != "" {
		panic("Environment variable must be provided: " + paramName)
	}

	//take all the required variables to work with the api
	var res *events.APIGatewayProxyResponse
	path := strings.Replace(request.RawPath, os.Getenv(ENV_URL_PREFIX), "", -1)
	method := request.RequestContext.HTTP.Method
	body := request.Body
	requestHeaders := request.Headers

	//build response headers
	responseHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	//call to handling the incoming request
	status, message := handlers.Handlers(path, method, body, requestHeaders, request)

	//build response
	res = &events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       string(message),
		Headers:    responseHeaders,
	}

	return res, nil
}

func ValidateEnvironmentVariables() (bool, string) {
	_, isParam := os.LookupEnv(ENV_SECRET_NAME)
	if !isParam {
		return isParam, ENV_SECRET_NAME
	}

	_, isParam = os.LookupEnv(ENV_URL_PREFIX)
	if !isParam {
		return isParam, ENV_URL_PREFIX
	}

	return isParam, ""
}
