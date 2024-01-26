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
	values := "VALUES ('" + util.ScapeString(product.ProdTitle) + "'"

	if len(product.ProdDescription) > 0 {
		sentence += ", Prod_Description"
		values += ", '" + util.ScapeString(product.ProdDescription) + "'"
	}

	if product.ProdPrice > 0 {
		sentence += ", Prod_Price"
		values += ", " + strconv.FormatFloat(product.ProdPrice, 'e', -1, 64)
	}

	if product.ProdCategoryId > 0 {
		sentence += ", Prod_CategoryId"
		values += ", " + strconv.Itoa(product.ProdCategoryId)
	}

	if product.ProdStock > 0 {
		sentence += ", Prod_Stock"
		values += ", " + strconv.Itoa(product.ProdStock)
	}

	if len(product.ProdPath) > 0 {
		sentence += ", Prod_Path"
		values += ", '" + util.ScapeString(product.ProdPath) + "'"
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

	response.ProdId, err = queryResult.LastInsertId()
	if err != nil {
		log.Default().Println("Error getting the last inserted id: " + err.Error())
		return response, err
	} else {
		response.ProdTitle = product.ProdTitle
		response.ProdDescription = product.ProdDescription
		response.ProdCreatedAt = time.Now()
	}

	log.Default().Printf("Product succesfully saved with id: %d.", response.ProdId)
	return response, nil
}

func UpdateProduct(prod models.ProductRequest) (models.ProductResponse, error) {
	log.Default().Println("Start to Update Product database")
	var err error
	var res models.ProductResponse

	err = DbConnect()
	if err != nil {
		log.Default().Println("Error connecting to the database..")
		return res, err
	}

	defer Db.Close()

	sentence := "UPDATE products SET "
	sentence = util.BuildUpdateSentence(sentence, "Prod_Title", prod.ProdTitle)
	sentence = util.BuildUpdateSentence(sentence, "Prod_CategoryId", prod.ProdCategoryId)
	sentence = util.BuildUpdateSentence(sentence, "Prod_Price", prod.ProdPrice)
	sentence = util.BuildUpdateSentence(sentence, "Prod_Stock", prod.ProdStock)
	sentence = util.BuildUpdateSentence(sentence, "Prod_Description", prod.ProdDescription)
	sentence = util.BuildUpdateSentence(sentence, "Prod_Path", prod.ProdPath)

	sentence += " WHERE Prod_Id = " + strconv.Itoa(int(prod.ProdId))
	log.Default().Println(sentence) //Only uncomment for debug purposes

	_, err = Db.Exec(sentence)
	if err != nil {
		log.Default().Println("Error executing update in the category table: " + err.Error())
		return res, err
	}
	res.ProdId = prod.ProdId

	log.Default().Printf("Product successfully updated with id: %d.", prod.ProdId)
	return res, nil
}
