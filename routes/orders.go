package routes

import (
	"encoding/json"
	"log"

	"github.com/SimonMora/bikesams_be/database"
	"github.com/SimonMora/bikesams_be/models"
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
