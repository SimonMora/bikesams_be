package database

import (
	"database/sql"
	"errors"
	"log"
	"strconv"

	"github.com/SimonMora/bikesams_be/models"
	"github.com/SimonMora/bikesams_be/util"
)

func InsertOrder(order models.OrderRequest, user string) (models.OrderResponse, error) {
	log.Default().Println("Start insert oreder process..")
	oResp := models.OrderResponse{}

	addExists, errA := AddressExists(user, order.OrderAddId)
	if !addExists {
		if errA != nil {
			return oResp, errA
		}
		return oResp, errors.New("Address Id doesn't exist in for the user: " + user + "in the database")
	}

	err := DbConnect()
	if err != nil {
		log.Default().Println("Error connecting to the database..")
		return oResp, err
	}
	defer Db.Close()

	sentenceOrder := "INSERT INTO orders (Order_UserUUID, Order_AddId, Order_Date, Order_Total) "
	sentenceOrder += "VALUES ('" + user + "', " + strconv.Itoa(order.OrderAddId) + ", '" + util.DateSqlFormat() + "', " + strconv.FormatFloat(order.OrderTotal, 'f', -1, 64) + ")"
	log.Default().Println(sentenceOrder) //Only uncomment for debug purposes

	var result sql.Result
	result, err = Db.Exec(sentenceOrder)
	if err != nil {
		log.Default().Println("Error inserting order in the database.. " + err.Error())
		return oResp, err
	}

	orderId, _ := result.LastInsertId()

	sentenceOD := "INSERT INTO orders_detail (OD_OrderId, OD_ProdId, OD_Quantity, OD_Price) VALUES "
	for index, od := range order.OrderDetails {
		log.Default().Println(index)
		sentenceOD += "(" + strconv.Itoa(int(orderId)) + ", " + strconv.Itoa(od.ODProdId) + ", " + strconv.Itoa(od.ODQuantity) + ", " + strconv.FormatFloat(od.ODPrice, 'f', -1, 64) + ")"
		if index != len(order.OrderDetails)-1 {
			sentenceOD += ", "
		}
	}
	log.Default().Println(sentenceOD) //Only uncomment for debug purposes

	result, err = Db.Exec(sentenceOD)
	if err != nil {
		log.Default().Println("Error inserting order details in the database.. " + err.Error())
		return oResp, err
	}

	oResp.FillOrderReq(order)
	oResp.OrderId = orderId
	oResp.OrderUserUUID = user

	log.Default().Println("Inserting order in the database was succesful..")
	return oResp, nil
}
