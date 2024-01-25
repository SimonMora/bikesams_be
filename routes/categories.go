package routes

import (
	"encoding/json"
	"log"

	"github.com/SimonMora/bikesams_be/database"
	"github.com/SimonMora/bikesams_be/models"
)

func ProcessCategoryRequest(body string, User string) (int, string) {
	log.Default().Println("Start to process category request..")
	var t models.Category

	// unmarshal body to struct and start body validations
	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		return 400, "Bad request, body is not category parseable."
	}

	if len(t.Categ_Name) == 0 {
		return 400, "Bad request, category name can't be empty."
	}

	if len(t.Categ_Path) == 0 {
		return 400, "Bad request, category path can't be empty."
	}

	//validate if the user is and admin or not
	isAdmin, msg := database.IsUserAdminValidate(User)
	if !isAdmin {
		return 400, msg
	}

	categoryProcessResult, insertError := database.InsertCategory(t)
	if insertError != nil {
		log.Default().Println("Error inserting the category with mesg: " + insertError.Error())
		return 400, insertError.Error()
	}

	catBytes, marshallErr := json.Marshal(categoryProcessResult)
	if marshallErr != nil {
		log.Default().Println("Error when transforming the category DTO to json: " + marshallErr.Error())
		return 400, marshallErr.Error()
	}

	return 200, string(catBytes)
}
