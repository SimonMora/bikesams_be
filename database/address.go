package database

import (
	"database/sql"
	"log"
	"strconv"
	"strings"

	"github.com/SimonMora/bikesams_be/models"
)

func InsertAddress(request models.AddressRequest, user string) (models.AddressResponse, error) {
	log.Default().Println("Start the Insert Address operation..")
	var add models.AddressResponse

	err := DbConnect()
	if err != nil {
		log.Default().Println("Error connecting to the database..")
		return add, err
	}
	defer Db.Close()

	sentence := "INSERT INTO addresses (Add_UserID,Add_Address,Add_City,Add_State,Add_PostalCode,Add_Phone,Add_Title,Add_Name)"
	sentence += "VALUES ('" + user + "', '" + request.AddAddress + "', '" + request.AddCity + "', '" + request.AddState + "', '" + request.AddPostalCode + "', '" + request.AddPhone + "', '" + request.AddTitle + "', '" + request.AddName + "')"
	//log.Default().Println(sentence) //Only uncomment for debug purposes

	var result sql.Result
	result, err = Db.Exec(sentence)
	if err != nil {
		log.Default().Println("Error executing the sentence in the database.. " + err.Error())
		return add, err
	}

	add.AddId, _ = result.LastInsertId()
	add.FillEntityReq(request)

	log.Default().Println("Insert Address successful..")
	return add, nil
}

func AddressExists(user string, id int) (bool, error) {
	log.Default().Println("Start verification Address existence..")

	err := DbConnect()
	if err != nil {
		log.Default().Println("Error connecting to the database..")
		return false, err
	}
	defer Db.Close()

	sentence := "SELECT 1 FROM addresses WHERE Add_Id = '" + strconv.Itoa(id) + "' AND Add_UserId = '" + user + "'"
	//log.Default().Println(sentence) //Only uncomment for debug purposes

	var rows *sql.Rows
	rows, err = Db.Query(sentence)
	if err != nil {
		log.Default().Println("Error executing the sentence in the database.. " + err.Error())
		return false, err
	}

	var result sql.NullString
	rows.Next()
	rows.Scan(&result)

	if result.String == "1" {
		return true, nil
	}

	log.Default().Println("Find Address successful..")
	return false, nil
}

func UpdateAddress(request models.AddressRequest) error {
	log.Default().Println("Start update Address processing..")

	err := DbConnect()
	if err != nil {
		log.Default().Println("Error connecting to the database..")
		return err
	}
	defer Db.Close()

	sentence := "UPDATE addresses SET "

	if request.AddTitle != "" {
		sentence += "Add_Title = '" + request.AddTitle + "', "
	}
	if request.AddName != "" {
		sentence += "Add_Name = '" + request.AddName + "', "
	}
	if request.AddAddress != "" {
		sentence += "Add_Address = '" + request.AddAddress + "', "
	}
	if request.AddCity != "" {
		sentence += "Add_City = '" + request.AddCity + "', "
	}
	if request.AddState != "" {
		sentence += "Add_State = '" + request.AddState + "', "
	}
	if request.AddPostalCode != "" {
		sentence += "Add_PostalCode = '" + request.AddPostalCode + "', "
	}
	if request.AddPhone != "" {
		sentence += "Add_Phone = '" + request.AddPhone + "', "
	}

	sentence, _ = strings.CutSuffix(sentence, ", ")
	sentence += "WHERE Add_Id = " + strconv.Itoa(request.AddId)
	//log.Default().Println(sentence) //Only uncomment for debug purposes

	_, err = Db.Query(sentence)
	if err != nil {
		log.Default().Println("Error executing the sentence in the database.. " + err.Error())
		return err
	}

	log.Default().Println("Update Address was successful..")
	return nil
}

func DeleteAddress(id int) error {
	log.Default().Println("Start delete Address processing..")

	err := DbConnect()
	if err != nil {
		log.Default().Println("Error connecting to the database..")
		return err
	}
	defer Db.Close()

	sentence := "DELETE FROM addresses WHERE Add_Id = " + strconv.Itoa(id)
	_, err = Db.Query(sentence)
	if err != nil {
		log.Default().Println("Error executing the sentence in the database.. " + err.Error())
		return err
	}

	log.Default().Println("Delete Address was successful..")
	return nil
}

func SelectAddressByUserId(user string) ([]models.AddressResponse, error) {
	var addResp []models.AddressResponse

	err := DbConnect()
	if err != nil {
		log.Default().Println("Error connecting to the database..")
		return addResp, err
	}
	defer Db.Close()

	sentence := "SELECT Add_Id, Add_Address, Add_City, Add_State, Add_PostalCode, Add_Phone, Add_Title, Add_Name FROM addresses"
	sentence += " WHERE Add_UserId = '" + user + "'"
	//log.Default().Println(sentence) //Only uncomment for debug purposes

	var rows *sql.Rows
	rows, err = Db.Query(sentence)
	if err != nil {
		log.Default().Println("Error executing the sentence in the database.. " + err.Error())
		return addResp, err
	}

	for rows.Next() {
		var addDb models.Address
		add := models.AddressResponse{}

		err = rows.Scan(
			&addDb.Add_Id, &addDb.Add_Address, &addDb.Add_City,
			&addDb.Add_State, &addDb.Add_PostalCode, &addDb.Add_Phone,
			&addDb.Add_Title, &addDb.Add_Name,
		)
		if err != nil {
			log.Default().Println("Error executing the sentence in the database.. " + err.Error())
			return []models.AddressResponse{}, err
		}

		add.FillEntityDb(addDb)
		addResp = append(addResp, add)
	}

	log.Default().Println("Addresses retrieved successfully..")
	return addResp, nil
}
