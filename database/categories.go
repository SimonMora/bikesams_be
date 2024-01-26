package database

import (
	"database/sql"
	"log"
	"strconv"
	"strings"

	"github.com/SimonMora/bikesams_be/models"
	"github.com/SimonMora/bikesams_be/util"
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

func UpdateCategory(category models.Category) error {
	log.Default().Println("Start to Update Category database")
	var err error

	err = DbConnect()
	if err != nil {
		log.Default().Println("Error connecting to the database..")
		return err
	}

	defer Db.Close()

	sentence := "UPDATE category SET "
	if category.Categ_Name != "" {
		sentence += "Categ_Name = '" + util.ScapeString(category.Categ_Name) + "'"
	}

	if category.Categ_Path != "" {
		if !strings.HasSuffix(sentence, "SET ") {
			sentence += ","
		}
		sentence += "Categ_Path = '" + util.ScapeString(category.Categ_Path) + "'"
	}

	sentence += " WHERE Categ_Id = " + strconv.Itoa(category.Categ_Id)
	//log.Default().Println(sentence) //Only uncomment for debug purposes

	_, err = Db.Exec(sentence)
	if err != nil {
		log.Default().Println("Error executing update in the category table: " + err.Error())
		return err
	}

	log.Default().Printf("Category successfully updated with id: %d.", category.Categ_Id)
	return nil
}

func DeleteCategory(id int) error {
	log.Default().Println("Start to Delete Category from database")
	var err error

	err = DbConnect()
	if err != nil {
		log.Default().Println("Error connecting to the database..")
		return err
	}
	defer Db.Close()

	sentence := "DELETE FROM category WHERE Categ_Id = " + strconv.Itoa(id)
	//log.Default().Println(sentence) //Only uncomment for debug purposes

	_, err = Db.Exec(sentence)
	if err != nil {
		log.Default().Println("Error executing update in the category table: " + err.Error())
		return err
	}

	log.Default().Printf("Category successfully deleted with id: %d.", id)
	return nil
}
