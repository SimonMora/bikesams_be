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
	User_Status    int16     `json:"userStatus"`
	User_DateAdd   time.Time `json:"userDateAdd"`
	User_DateUpg   time.Time `json:"userDateUpg"`
}

type UserResponse struct {
	UserUUID       string    `json:"userUUID"`
	User_Email     string    `json:"userEmail"`
	User_FirstName string    `json:"userFirstName"`
	User_LastName  string    `json:"userLastName"`
	User_Status    int16     `json:"userStatus"`
	User_DateAdd   time.Time `json:"userDateAdd"`
	User_DateUpg   time.Time `json:"userDateUpg,omitempty"`
}

type User struct {
	UserUUID       sql.NullString `json:"userUUID"`
	User_Email     sql.NullString `json:"userEmail"`
	User_FirstName sql.NullString `json:"userFirstName"`
	User_LastName  sql.NullString `json:"userLastName"`
	User_Status    sql.NullInt16  `json:"userStatus"`
	User_DateAdd   sql.NullTime   `json:"userDateAdd"`
	User_DateUpg   sql.NullTime   `json:"userDateUpg"`
}

func (User *UserResponse) ParseUser(dbUser User) {
	User.UserUUID = dbUser.UserUUID.String
	User.User_Email = dbUser.User_Email.String
	User.User_FirstName = dbUser.User_FirstName.String
	User.User_LastName = dbUser.User_LastName.String
	User.User_Status = dbUser.User_Status.Int16
	User.User_DateAdd = dbUser.User_DateAdd.Time
	User.User_DateUpg = dbUser.User_DateUpg.Time
}
