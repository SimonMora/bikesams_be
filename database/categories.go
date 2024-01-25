package database

import (
	"database/sql"
	"log"

	"github.com/SimonMora/bikesams_be/models"
	_ "github.com/go-sql-driver/mysql"
)

func InsertCategory(category models.Category) (models.CategoryProcessResult, error) {
	log.Default().Println("Start to insert Category into database")
	var response models.CategoryProcessResult

	err := DbConnect()
	if err != nil {
		log.Default().Println("Error connecting to the database..")
		return response, err
	}

	defer Db.Close()

	sentence := "INSERT INTO category (Categ_Name, Categ_Path) VALUES ('" + category.Categ_Name + "', '" + category.Categ_Path + "')"
	//log.Default().Println(sentence) //Only uncomment for debug purposes

	var queryResult sql.Result

	queryResult, err = Db.Exec(sentence)
	if err != nil {
		log.Default().Println("Error executing insert in the category table: " + err.Error())
		return response, err
	}

	response.CategId, err = queryResult.LastInsertId()
	if err != nil {
		log.Default().Println("Error getting the last inserted id: " + err.Error())
		return response, err
	}

	log.Default().Printf("Category succesfully saved with id: %d.", response.CategId)
	return response, nil

}
