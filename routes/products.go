package routes

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"

	"github.com/SimonMora/bikesams_be/database"
	"github.com/SimonMora/bikesams_be/models"
	"github.com/aws/aws-lambda-go/events"
)

func InsertProducts(body string, User string) (int, string) {
	log.Default().Println("Start Insert products process.")
	var response models.ProductResponse
	var product models.Product
	var err error

	err = json.Unmarshal([]byte(body), &product)
	if err != nil {
		log.Default().Println("Error parsing the body to product: " + err.Error())
		return 400, "Error parsing the body to product, please verify"
	}

	if len(product.ProdTitle) == 0 {
		return 400, "Product title must be specified"
	}

	//validate if the user is and admin or not
	isAdmin, msg := database.IsUserAdminValidate(User)
	if !isAdmin {
		return 400, msg
	}

	response, err = database.InsertProducts(product)
	if err != nil {
		log.Default().Println("Error inserting the product: " + err.Error())
		return 500, "Error parsing inserting the product: " + err.Error()
	}

	byteRes, errM := json.Marshal(response)
	if errM != nil {
		log.Default().Println("Error parsing the database return: " + err.Error())
		return 500, err.Error()
	}

	log.Default().Println("Product insert was successful, with product id: " + strconv.Itoa(int(response.ProdId)))
	return 200, string(byteRes)

}

func UpdateProduct(body string, User string, id int) (int, string) {
	log.Default().Println("Start to process to update category..")
	if id == 0 {
		return 400, "Product id is required to update the product."
	}

	var prod models.ProductRequest

	// unmarshal body to struct and start body validations
	err := json.Unmarshal([]byte(body), &prod)
	if err != nil {
		return 400, "Bad request, body is not product parseable."
	}

	//validate if the user is and admin or not
	isAdmin, msg := database.IsUserAdminValidate(User)
	if !isAdmin {
		return 400, msg
	}

	prod.ProdId = id
	_, updateErr := database.UpdateProduct(prod)
	if updateErr != nil {
		return 400, "Error when trying to update the product: " + strconv.Itoa(id) + " > " + updateErr.Error()
	}

	return 200, "Updated entity"
}

func DeleteProduct(User string, id int) (int, string) {
	if id == 0 {
		return 400, "Product id is required to delete products."
	}

	//validate if the user is and admin or not
	isAdmin, msg := database.IsUserAdminValidate(User)
	if !isAdmin {
		return 400, msg
	}

	err := database.DeleteProduct(id)
	if err != nil {
		return 400, "Error when trying to delete the product: " + strconv.Itoa(id) + " > " + err.Error()
	}

	return 200, "Product deleted"
}

func SelectProduct(request events.APIGatewayV2HTTPRequest) (int, string) {
	var t models.ProductRequest
	var res models.PaginatedProduct
	var page, pageSize int
	var orderType, orderField string
	var err error

	param := request.QueryStringParameters

	page, _ = strconv.Atoi(param["page"])
	pageSize, _ = strconv.Atoi(param["pageSize"])
	orderType = param["orderType"]
	orderField = param["orderField"]

	if !strings.Contains("ITDFPCS", orderType) {
		orderField = ""
	}

	var choice string
	if len(param["prodId"]) > 0 {
		choice = "P"
		t.ProdId, _ = strconv.Atoi(param["prodId"])
	}
	if len(param["search"]) > 0 {
		choice = "S"
		t.ProdSearch = param["search"]
	}
	if len(param["categId"]) > 0 {
		choice = "C"
		t.ProdCategoryId, _ = strconv.Atoi(param["categId"])
	}
	if len(param["slug"]) > 0 {
		choice = "U"
		t.ProdPath = param["slug"]
	}
	if len(param["slugCateg"]) > 0 {
		choice = "K"
		t.ProdCategPath = param["slugCateg"]
	}

	log.Default().Println(param)

	res, err = database.SelectProduct(t, choice, page, pageSize, orderType, orderField)
	if err != nil {
		log.Default().Println("There was an error trying to get the products from the database, search type: " + choice + " in product database. > " + err.Error())
		return 500, "There was an error trying to get the products from the database, search type: " + choice + " in product database."
	}

	byteResp, errM := json.Marshal(res)
	if errM != nil {
		log.Default().Println("There was an error unmarshal product. " + errM.Error())
		return 500, "There was an error unmarshal product."
	}

	return 200, string(byteResp)
}
