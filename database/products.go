package database

import (
	"database/sql"
	"errors"
	"log"
	"strconv"
	"strings"
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
	//log.Default().Println(sentence) //Only uncomment for debug purposes

	_, err = Db.Exec(sentence)
	if err != nil {
		log.Default().Println("Error executing update in the category table: " + err.Error())
		return res, err
	}
	res.ProdId = int64(prod.ProdId)

	log.Default().Printf("Product successfully updated with id: %d.", prod.ProdId)
	return res, nil
}

func DeleteProduct(id int) error {
	log.Default().Println("Start to Delete Products from database")
	var err error

	err = DbConnect()
	if err != nil {
		log.Default().Println("Error connecting to the database..")
		return err
	}
	defer Db.Close()

	sentence := "DELETE FROM products WHERE Prod_Id = " + strconv.Itoa(id)
	//log.Default().Println(sentence) //Only uncomment for debug purposes

	_, err = Db.Exec(sentence)
	if err != nil {
		log.Default().Println("Error executing update in the products table: " + err.Error())
		return err
	}

	log.Default().Printf("Category successfully deleted with id: %d.", id)
	return nil
}

func SelectProduct(prod models.ProductRequest, choice string, page int, pageSize int, orderType string, orderField string) (models.PaginatedProduct, error) {
	log.Default().Println("Start to Retrieve Products from database..")
	var Resp models.PaginatedProduct
	var Prods []models.Product

	err = DbConnect()
	if err != nil {
		log.Default().Println("Error connecting to the database..")
		return Resp, err
	}
	defer Db.Close()

	var sentence, countSentence, where, limit string

	sentence = "SELECT p.Prod_Id, p.Prod_Title, p.Prod_Description, p.Prod_CreatedAt, p.Prod_Updated, p.Prod_Price, p.Prod_Path, p.Prod_CategoryId, p.Prod_Stock FROM products p"
	countSentence = "SELECT count(*) as records FROM products p"

	switch choice {
	case "P":
		where = " WHERE p.Prod_Id = " + strconv.Itoa(prod.ProdId)
	case "S":
		where = " WHERE UCASE(CONCAT(p.Prod_Description, p.Prod_Title)) like '%" + strings.ToUpper(util.ScapeString(prod.ProdSearch)) + "%'"
	case "C":
		where = " WHERE p.Prod_CategoryId = " + strconv.Itoa(prod.ProdCategoryId)
	case "U":
		where = " WHERE UCASE(p.Prod_Path) like '%" + strings.ToUpper(util.ScapeString(prod.ProdPath)) + "%'"
	case "K":
		where = "INNER JOIN category c ON p.Prod_CategoryId = c.Categ_Id WHERE c.Categ_Path LIKE '%" + strings.ToUpper(util.ScapeString(prod.ProdCategPath)) + "%'"
	}

	countSentence += where
	log.Default().Println(countSentence) //Only uncomment for debug purposes

	var result *sql.Rows

	result, err = Db.Query(countSentence)
	if err != nil {
		log.Default().Println("Error executing count select in the products table: " + err.Error())
		return Resp, err
	}

	result.Next()
	var reco sql.NullInt32

	err := result.Scan(&reco)
	if err != nil {
		log.Default().Println("Error when extracting the result count of product records: " + err.Error())
		return Resp, err
	}

	records := int(reco.Int32)

	result.Close()

	if page > 0 {
		if records > pageSize {
			limit = " LIMIT " + strconv.Itoa(pageSize)
			if page > 1 {
				offset := pageSize * (page - 1)
				limit += " OFFSET " + strconv.Itoa(offset)
			}
		} else {
			limit = ""
		}
	}

	//ITDFPCS
	var orderBy string
	if len(orderField) > 0 {
		switch orderField {
		case "I":
			orderBy = " ORDER BY Prod_Id "
		case "T":
			orderBy = " ORDER BY Prod_Title "
		case "D":
			orderBy = " ORDER BY Prod_Description "
		case "F":
			orderBy = " ORDER BY Prod_CreatedAt "
		case "P":
			orderBy = " ORDER BY Prod_Price "
		case "C":
			orderBy = " ORDER BY Prod_CategoryId "
		case "S":
			orderBy = " ORDER BY Prod_Stock "
		}

		if orderType == "D" {
			orderBy += " DESC"
		}
	}

	sentence += where + orderBy + limit
	log.Default().Println(sentence) //Only uncomment for debug purposes

	result, err = Db.Query(sentence)
	if err != nil {
		log.Default().Println("Error executing main select in the products table: " + err.Error())
		return Resp, err
	}

	for result.Next() {
		var p models.Product
		var prodId sql.NullInt32
		var prodTitle sql.NullString
		var prodDescription sql.NullString
		var prodCreatedAt sql.NullTime
		var prodUpdated sql.NullTime
		var prodPrice sql.NullFloat64
		var prodPath sql.NullString
		var prodCategoryId sql.NullInt32
		var prodStock sql.NullInt32

		err := result.Scan(&prodId, &prodTitle, &prodDescription, &prodCreatedAt, &prodUpdated, &prodPrice, &prodPath, &prodCategoryId, &prodStock)
		if err != nil {
			log.Default().Println("Error parsing the products retrieved from table: " + err.Error())
			return Resp, err
		}

		p.ProdId = int64(prodId.Int32)
		p.ProdTitle = prodTitle.String
		p.ProdDescription = prodDescription.String
		p.ProdCreatedAt = prodCreatedAt.Time
		p.ProdUpdated = prodUpdated.Time
		p.ProdPrice = prodPrice.Float64
		p.ProdPath = prodPath.String
		p.ProdCategoryId = int(prodCategoryId.Int32)
		p.ProdStock = int(prodStock.Int32)

		Prods = append(Prods, p)
	}

	Resp.TotalItems = records
	Resp.Data = Prods

	log.Default().Println("Products retrieved from the database..")
	return Resp, nil
}

func UpdateStock(prod models.ProductRequest) error {
	log.Default().Println("Start to Update Products Stock in the database..")
	if prod.ProdStock == 0 {
		return errors.New("[ERROR] the prod stock is required to update the stock.")
	}

	err := DbConnect()
	if err != nil {
		log.Default().Println("Error connecting to the database..")
		return err
	}
	defer Db.Close()

	sentence := "UPDATE products SET Prod_Stock = Prod_Stock + " + strconv.Itoa(prod.ProdStock) + " WHERE Prod_Id = " + strconv.Itoa(prod.ProdId)

	_, err = Db.Exec(sentence)
	if err != nil {
		log.Default().Println("Error executing update in the category table: " + err.Error())
		return err
	}

	log.Default().Println("Products stock updated in the database, product id : " + strconv.Itoa(prod.ProdId))
	return nil
}
