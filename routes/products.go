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

	prod.ProdId = int64(id)
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
