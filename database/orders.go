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
	log.Default().Println("Start insert order process..")
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
	//log.Default().Println(sentenceOrder) //Only uncomment for debug purposes

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
	//log.Default().Println(sentenceOD) //Only uncomment for debug purposes

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

func SelectOrders(user string, dateFrom string, dateTo string, page int, orderId int) ([]models.OrderResponse, error) {
	log.Default().Println("Start select order process..")
	var orders []models.OrderResponse

	sentence := "SELECT Order_Id, Order_UserUUID, Order_AddId, Order_Date, Order_Total FROM orders "

	if orderId != 0 {
		sentence += "WHERE Order_Id = " + strconv.Itoa(orderId)
	} else {
		offset := 0

		if page == 0 {
			page = 1
		}

		if page > 1 {
			offset = (page * 10) - 10
		}

		if len(dateFrom) != 0 {
			if len(dateFrom) == 10 {
				dateFrom += "T00:00:00"
			}
			if len(dateTo) == 10 {
				dateTo += "T23:59:59"
			}
			sentence += "WHERE Order_Date BETWEEN '" + dateFrom + "' AND '" + dateTo + "' " + "AND Order_UserUUID = '" + user + "'"
		} else {
			sentence += "WHERE Order_UserUUID = '" + user + "'"
		}

		if offset != 0 {
			sentence += " OFFSET " + strconv.Itoa(offset)
		}

		sentence += " LIMIT 10"
	}

	log.Default().Println(sentence) //Only uncomment for debug purposes

	err := DbConnect()
	if err != nil {
		log.Default().Println("Error connecting to the database..")
		return orders, err
	}
	defer Db.Close()

	var rows *sql.Rows
	rows, err = Db.Query(sentence)
	if err != nil {
		log.Default().Println("Error executing the sentence in the database.. " + err.Error())
		return orders, err
	}
	defer rows.Close()

	for rows.Next() {
		var o models.Order
		orderResp := models.OrderResponse{}
		errS := rows.Scan(&o.Order_Id, &o.Order_UserUUID, &o.Order_AddId, &o.Order_Date, &o.Order_Total)
		if errS != nil {
			log.Default().Println("Error scanning the order rows in the database.. " + errS.Error())
			return []models.OrderResponse{}, errS
		}
		orderResp.FillOrderDb(o)

		sentenceD := "SELECT OD_Id, OD_OrderId, OD_ProdId, OD_Quantity, OD_Price FROM orders_detail WHERE OD_OrderId = " + strconv.Itoa(int(orderResp.OrderId))
		log.Default().Println(sentenceD) //Only uncomment for debug purposes

		var rowsD *sql.Rows
		rowsD, errS = Db.Query(sentenceD)
		if errS != nil {
			log.Default().Println("Error executing the details sentence in the database.. " + errS.Error())
			return []models.OrderResponse{}, errS
		}

		for rowsD.Next() {
			var od models.OrderDetails
			orderDResp := models.OrderDetailsResponse{}

			err := rowsD.Scan(&od.OD_Id, &od.OD_OderId, &od.OD_ProdId, &od.OD_Quantity, &od.OD_Price)
			if err != nil {
				log.Default().Println("Error scanning the order details rows in the database.. " + err.Error())
				return []models.OrderResponse{}, err
			}

			orderDResp.FillOrderDetailDb(od)
			orderResp.OrderDetails = append(orderResp.OrderDetails, orderDResp)
		}

		orders = append(orders, orderResp)
	}

	log.Default().Println("Select orders from the database was succesful..")
	return orders, nil
}
