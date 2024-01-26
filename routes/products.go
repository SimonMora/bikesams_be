package routes

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/SimonMora/bikesams_be/database"
	"github.com/SimonMora/bikesams_be/models"
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

	if len(product.Prod_Title) == 0 {
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

	log.Default().Println("Product insert was successful, with product id: " + strconv.Itoa(int(response.Prod_Id)))
	return 200, string(byteRes)

}
