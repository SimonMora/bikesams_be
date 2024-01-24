package handlers

import (
	"log"
	//"strconv"

	"github.com/aws/aws-lambda-go/events"
)

func Handlers(path string, method string, body string, headers map[string]string, request *events.APIGatewayV2HTTPRequest) (int, string) {
	log.Default().Printf("Start to handle the request from: %s with method type %s\n", path, method)

	/*id := request.PathParameters["id"]
	numId, _ := strconv.Atoi(id)*/

	//Missing handlers actions

	return 400, "Invalid Method"
}
