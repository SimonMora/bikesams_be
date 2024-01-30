package routes

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/SimonMora/bikesams_be/database"
	"github.com/SimonMora/bikesams_be/models"
)

func InsertAddress(body string, user string) (int, string) {
	var add models.AddressRequest
	var addResp models.AddressResponse

	err := json.Unmarshal([]byte(body), &add)
	if err != nil {
		log.Default().Println("Error unmarshaling body to Address class: " + err.Error())
		return 400, "Body is not address parseable"
	}

	if add.AddTitle == "" {
		return 400, "Address Title is required"
	}
	if add.AddName == "" {
		return 400, "Address Name is required"
	}
	if add.AddAddress == "" {
		return 400, "Address Address is required"
	}
	if add.AddCity == "" {
		return 400, "Address City is required"
	}
	if add.AddState == "" {
		return 400, "Address State is required"
	}
	if add.AddPostalCode == "" {
		return 400, "Address PostalCode is required"
	}
	if add.AddPhone == "" {
		return 400, "Address Phone is required"
	}

	addResp, err = database.InsertAddress(add, user)
	if err != nil {
		log.Default().Println("Error on insert the Address in database: " + err.Error())
		return 500, "Database access error: " + err.Error()
	}

	addBytes, errU := json.Marshal(addResp)
	if errU != nil {
		log.Default().Println("Error on marshaling the Address inserted in database: " + err.Error())
		return 500, "Marshaling error: " + err.Error()
	}

	return 200, string(addBytes)
}

func UpdateAddress(body string, user string, addId int) (int, string) {
	var add models.AddressRequest

	err := json.Unmarshal([]byte(body), &add)
	if err != nil {
		log.Default().Println("Error unmarshaling body to Address class: " + err.Error())
		return 400, "Body is not address parseable"
	}

	exist, errE := database.AddressExists(user, addId)
	if !exist {
		if errE != nil {
			log.Default().Println("The was an error searching Address.. " + errE.Error())
			return 400, errE.Error()
		}
		log.Default().Println("The Address was not found for user: " + user + ", address id: " + strconv.Itoa(addId))
		return 400, "Address was not found"
	}

	add.AddId = addId
	err = database.UpdateAddress(add)
	if err != nil {
		log.Default().Println("The Address update failed: " + err.Error())
		return 500, "Error on update" + err.Error()
	}

	return 200, "Updated entity"
}

func DeleteAddress(user string, id int) (int, string) {
	exist, err := database.AddressExists(user, id)
	if !exist {
		if err != nil {
			log.Default().Println("The was an error searching Address.. " + err.Error())
			return 400, err.Error()
		}
		log.Default().Println("The Address was not found for user: " + user + ", address id: " + strconv.Itoa(id))
		return 400, "Address was not found"
	}

	err = database.DeleteAddress(id)
	if err != nil {
		log.Default().Println("The Address delete failed: " + err.Error())
		return 500, "Error on update" + err.Error()
	}

	return 200, "Deleted entity"
}

func SelectAddress(user string) (int, string) {
	addResp, err := database.SelectAddressByUserId(user)
	if err != nil {
		log.Default().Println("Error retrieving the Address from the database: " + err.Error())
		return 500, "Database access error: " + err.Error()
	}

	addBytes, errU := json.Marshal(addResp)
	if errU != nil {
		log.Default().Println("Error on marshaling the Address retrieved from database: " + err.Error())
		return 500, "Marshaling error: " + err.Error()
	}

	return 200, string(addBytes)
}
