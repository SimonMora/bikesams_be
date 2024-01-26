package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/SimonMora/bikesams_be/auth"
	"github.com/SimonMora/bikesams_be/routes"
	"github.com/aws/aws-lambda-go/events"
)

const CATEGORY = "cate"
const PRODUCTS = "prod"
const USERS = "user"
const ADDRESS = "addr"
const ORDERS = "orde"
const STOCK = "stoc"

func Handlers(path string, method string, body string, headers map[string]string, request events.APIGatewayV2HTTPRequest) (int, string) {
	log.Default().Printf("Start to handle the request from: %s with method type %s\n", path, method)

	id := request.PathParameters["id"]
	numId, _ := strconv.Atoi(id)

	isOk, status, userMsg := validateAuthorization(path, method, headers)
	if !isOk {
		log.Default().Println("Error with token validation: " + userMsg)
		return status, userMsg
	}

	switch path[0:4] {
	case CATEGORY:
		return handleCategoriesRequest(body, path, method, userMsg, numId, request)
	case PRODUCTS:
		return handleProductsRequest(body, path, method, userMsg, numId, request)
	case USERS:
		return handleUsersRequest(body, path, method, userMsg, id, request)
	case ADDRESS:
		return handleAddressRequest(body, path, method, userMsg, numId, request)
	case ORDERS:
		return handleOrdersRequest(body, path, method, userMsg, numId, request)
	case STOCK:
		return handleStockRequest(body, path, method, userMsg, numId, request)
	}

	return 400, "Invalid Method"
}

func validateAuthorization(path string, method string, headers map[string]string) (bool, int, string) {
	log.Default().Println("Start to validate Authorization..")

	if (path == "products" && method == http.MethodGet) ||
		(path == "category" && method == http.MethodGet) {
		log.Default().Println("Authorization is not required..")
		return true, 200, ""
	}

	token := headers["authorization"]
	if len(token) == 0 {
		return false, 401, "No token provided."
	}

	isOK, err, msg := auth.TokenValidate(token)
	if !isOK {
		if err != nil {
			log.Default().Println("Error with token validation: " + err.Error())
			return false, 401, err.Error()
		} else {
			log.Default().Println("Error with token validation: " + msg)
			return false, 401, msg
		}
	}

	log.Default().Println("The token is OK..")
	return true, 200, msg

}

func handleCategoriesRequest(body string, path string, method string, user string, id int, event events.APIGatewayV2HTTPRequest) (int, string) {
	switch method {
	case http.MethodPost:
		return routes.ProcessCategoryRequest(body, user)
	case http.MethodPut:
		return routes.UpdateCategory(body, user, id)
	}
	return 400, "Invalid Method"
}

func handleProductsRequest(body string, path string, method string, user string, id int, event events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Invalid Method"
}

func handleUsersRequest(body string, path string, method string, user string, id string, event events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Invalid Method"
}

func handleAddressRequest(body string, path string, method string, user string, id int, event events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Invalid Method"
}

func handleOrdersRequest(body string, path string, method string, user string, id int, event events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Invalid Method"
}

func handleStockRequest(body string, path string, method string, user string, id int, event events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Invalid Method"
}
