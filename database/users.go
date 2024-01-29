package database

import (
	"database/sql"
	"log"
	"strings"

	"github.com/SimonMora/bikesams_be/models"
	"github.com/SimonMora/bikesams_be/util"
	_ "github.com/go-sql-driver/mysql"
)

func SearchUser(user string) (models.UserResponse, bool) {
	log.Default().Println("Start User Search for user with id: " + user)
	var userDb models.User
	var userResp models.UserResponse
	var result *sql.Rows

	err := DbConnect()
	if err != nil {
		log.Default().Println("Error connecting to the database..")
		return userResp, false
	}
	defer Db.Close()

	sentence := "SELECT User_UUID, User_Email, User_FirstName, User_LastName FROM users WHERE User_UUID = '" + user + "'"

	result, err = Db.Query(sentence)
	if err != nil {
		log.Default().Println("Error retrieving user from database.. " + err.Error())
		return userResp, false
	}

	result.Next()
	err = result.Scan(&userDb.UserUUID, &userDb.User_Email, &userDb.User_FirstName, &userDb.User_LastName)
	if err != nil {
		log.Default().Println("Error parsing the user retrieved from the database.. " + err.Error())
		return userResp, false
	} else {
		userResp.UserUUID = userDb.UserUUID.String
		userResp.User_Email = userDb.User_Email.String
		userResp.User_FirstName = userDb.User_FirstName.String
		userResp.User_LastName = userDb.User_LastName.String
	}

	return userResp, true
}

func UpdateUser(request models.UserRequest, user string) (models.UserResponse, error) {
	log.Default().Println("Start User Update for user with id: " + user)
	var userResp models.UserResponse

	err := DbConnect()
	if err != nil {
		log.Default().Println("Error connecting to the database..")
		return userResp, err
	}

	defer Db.Close()

	sentence := "UPDATE users SET "
	if len(request.User_FirstName) > 0 {
		sentence += " User_FirstName = '" + util.ScapeString(request.User_FirstName) + "' "
	}

	if len(request.User_LastName) > 0 {
		if !strings.HasSuffix(sentence, "SET ") {
			sentence += ", "
		}
		sentence += " User_LastName = '" + util.ScapeString(request.User_LastName) + "' "
	}

	sentence += ", User_DateUpg = '" + util.DateSqlFormat() + "' WHERE User_UUID = '" + user + "'"
	log.Default().Println(sentence) //Only uncomment for debug purposes

	_, err = Db.Exec(sentence)
	if err != nil {
		log.Default().Println("Error executing sentence in the database.." + err.Error())
		return userResp, err
	}
	userResp, _ = SearchUser(user)

	log.Default().Println("User Update sucessfull..")
	return userResp, nil
}
