package database

import (
	"database/sql"
	"log"
	"strconv"

	//"strings"
	"time"

	"github.com/SimonMora/bikesams_be/models"
	"github.com/SimonMora/bikesams_be/util"
	_ "github.com/go-sql-driver/mysql"
)

func InsertProducts(product models.Product) (models.ProductResponse, error) {

	log.Default().Println("Start to insert Product into database")
	var response models.ProductResponse

	err := DbConnect()
	if err != nil {
		log.Default().Println("Error connecting to the database..")
		return response, err
	}
	defer Db.Close()

	sentence := "INSERT INTO products(Prod_Title"
	values := "VALUES ('" + util.ScapeString(product.Prod_Title) + "'"

	if len(product.Prod_Description) > 0 {
		sentence += ", Prod_Description"
		values += ", '" + util.ScapeString(product.Prod_Description) + "'"
	}

	if product.Prod_Price > 0 {
		sentence += ", Prod_Price"
		values += ", " + strconv.FormatFloat(product.Prod_Price, 'e', -1, 64)
	}

	if product.Prod_CategoryId > 0 {
		sentence += ", Prod_CategoryId"
		values += ", " + strconv.Itoa(product.Prod_CategoryId)
	}

	if product.Prod_Stock > 0 {
		sentence += ", Prod_Stock"
		values += ", " + strconv.Itoa(product.Prod_Stock)
	}

	if len(product.Prod_Path) > 0 {
		sentence += ", Prod_Path"
		values += ", '" + util.ScapeString(product.Prod_Path) + "'"
	}

	sentence += ")"
	values += ")"

	sentence += values
	//log.Default().Println(sentence) //Only uncomment for debug purposes

	var queryResult sql.Result

	queryResult, err = Db.Exec(sentence)
	if err != nil {
		log.Default().Println("Error executing insert in the category table: " + err.Error())
		return response, err
	}

	response.Prod_Id, err = queryResult.LastInsertId()
	if err != nil {
		log.Default().Println("Error getting the last inserted id: " + err.Error())
		return response, err
	} else {
		response.Prod_Title = product.Prod_Title
		response.Prod_Description = product.Prod_Description
		response.Prod_CreatedAt = time.Now()
	}

	log.Default().Printf("Product succesfully saved with id: %d.", response.Prod_Id)
	return response, nil
}
