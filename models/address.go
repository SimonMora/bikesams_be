package models

import "database/sql"

type Address struct {
	Add_Id         sql.NullInt64
	Add_Title      sql.NullString
	Add_Name       sql.NullString
	Add_Address    sql.NullString
	Add_City       sql.NullString
	Add_State      sql.NullString
	Add_PostalCode sql.NullString
	Add_Phone      sql.NullString
}

type AddressRequest struct {
	AddId         int    `json:"addId"`
	AddTitle      string `json:"addTitle"`
	AddName       string `json:"addName"`
	AddAddress    string `json:"addAddress"`
	AddCity       string `json:"addCity"`
	AddState      string `json:"addState"`
	AddPostalCode string `json:"addPostalCode"`
	AddPhone      string `json:"addPhone"`
}

type AddressResponse struct {
	AddId         int64  `json:"addId"`
	AddTitle      string `json:"addTitle"`
	AddName       string `json:"addName"`
	AddAddress    string `json:"addAddress"`
	AddCity       string `json:"addCity"`
	AddState      string `json:"addState"`
	AddPostalCode string `json:"addPostalCode"`
	AddPhone      string `json:"addPhone"`
}

func (add *AddressResponse) FillEntityReq(addReq AddressRequest) {
	add.AddAddress = addReq.AddAddress
	add.AddCity = addReq.AddCity
	add.AddName = addReq.AddName
	add.AddPhone = addReq.AddPhone
	add.AddPostalCode = addReq.AddPostalCode
	add.AddState = addReq.AddState
	add.AddTitle = addReq.AddTitle
}

func (add *AddressResponse) FillEntityDb(addDb Address) {
	add.AddAddress = addDb.Add_Address.String
	add.AddCity = addDb.Add_City.String
	add.AddName = addDb.Add_Name.String
	add.AddPhone = addDb.Add_Phone.String
	add.AddPostalCode = addDb.Add_PostalCode.String
	add.AddState = addDb.Add_State.String
	add.AddTitle = addDb.Add_Title.String
}
