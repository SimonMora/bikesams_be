package routes

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/SimonMora/bikesams_be/database"
	"github.com/SimonMora/bikesams_be/models"
	"github.com/aws/aws-lambda-go/events"
)

func InsertOrder(body string, user string) (int, string) {
	var o models.OrderRequest
	var oResp models.OrderResponse

	err := json.Unmarshal([]byte(body), &o)
	if err != nil {
		log.Default().Println("Error parsing the body to order: " + err.Error())
		return 400, "Body is not order parseable"
	}

	isOk, msg := validOrder(o)
	if !isOk {
		return 400, msg
	}

	oResp, err = database.InsertOrder(o, user)
	if err != nil {
		log.Default().Println("Error inserting order in the database: " + err.Error())
		return 500, "Error inserting order" + err.Error()
	}

	obytes, errM := json.Marshal(oResp)
	if errM != nil {
		log.Default().Println("Error parsing order inserted in the database: " + err.Error())
		return 500, "Error parsing order" + err.Error()
	}

	return 200, string(obytes)
}

func validOrder(order models.OrderRequest) (bool, string) {
	if order.OrderTotal == 0 {
		return false, "Order Total must be specified"
	}

	if order.OrderAddId == 0 {
		return false, "Order Address must be specified"
	}

	var count int16
	for _, od := range order.OrderDetails {
		if od.ODProdId == 0 {
			return false, "Product Id must be specified"
		}
		if od.ODQuantity == 0 {
			return false, "Product Quantity must be specified"
		}

		count++
	}

	if count == 0 {
		return false, "Must specified items in the order"
	}

	return true, ""
}

func SelectOrders(user string, request events.APIGatewayV2HTTPRequest) (int, string) {
	var dateFrom, dateTo string
	var page int
	var orderId int

	log.Default().Println(request.QueryStringParameters)

	if len(request.QueryStringParameters["dateFrom"]) != 0 {
		dateFrom = request.QueryStringParameters["dateFrom"]
		if len(request.QueryStringParameters["dateTo"]) != 0 {
			dateTo = request.QueryStringParameters["dateTo"]
		} else {
			return 400, "Date end must be provided if Date from was provided"
		}
	}

	if len(request.QueryStringParameters["page"]) != 0 {
		page, _ = strconv.Atoi(request.QueryStringParameters["page"])
	}
	if len(request.QueryStringParameters["orderId"]) != 0 {
		orderId, _ = strconv.Atoi(request.QueryStringParameters["orderId"])
	}

	result, err := database.SelectOrders(user, dateFrom, dateTo, page, orderId)
	if err != nil {
		log.Default().Println("Error retrieving orders from database: " + err.Error())
		return 500, "Error retrieving orders: " + err.Error()
	}

	oBytes, errM := json.Marshal(result)
	if errM != nil {
		log.Default().Println("Error parsing orders from database: " + errM.Error())
		return 500, "Error parsing orders: " + err.Error()
	}

	return 200, string(oBytes)
}
