package models

import (
	"database/sql"
	"time"
)

type UserRequest struct {
	UserUUID       string    `json:"userUUID"`
	User_Email     string    `json:"userEmail"`
	User_FirstName string    `json:"userFirstName"`
	User_LastName  string    `json:"userLastName"`
	User_Status    string    `json:"userStatus"`
	User_DateAdd   time.Time `json:"userDateAdd"`
	User_DateUpg   time.Time `json:"userDateUpg"`
}

type UserResponse struct {
	UserUUID       string `json:"userUUID"`
	User_Email     string `json:"userEmail"`
	User_FirstName string `json:"userFirstName"`
	User_LastName  string `json:"userLastName"`
}

type User struct {
	UserUUID       sql.NullString `json:"userUUID"`
	User_Email     sql.NullString `json:"userEmail"`
	User_FirstName sql.NullString `json:"userFirstName"`
	User_LastName  sql.NullString `json:"userLastName"`
	User_Status    sql.NullString `json:"userStatus"`
	User_DateAdd   sql.NullTime   `json:"userDateAdd"`
	User_DateUpg   sql.NullTime   `json:"userDateUpg"`
}
