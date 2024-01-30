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

	sentence := "SELECT User_UUID, User_Email, User_FirstName, User_LastName, User_DateUpg FROM users WHERE User_UUID = '" + user + "'"

	result, err = Db.Query(sentence)
	if err != nil {
		log.Default().Println("Error retrieving user from database.. " + err.Error())
		return userResp, false
	}

	result.Next()
	err = result.Scan(&userDb.UserUUID, &userDb.User_Email, &userDb.User_FirstName, &userDb.User_LastName, &userDb.User_DateUpg)
	if err != nil {
		log.Default().Println("Error parsing the user retrieved from the database.. " + err.Error())
		return userResp, false
	} else {
		userResp.UserUUID = userDb.UserUUID.String
		userResp.User_Email = userDb.User_Email.String
		userResp.User_FirstName = userDb.User_FirstName.String
		userResp.User_LastName = userDb.User_LastName.String
		userResp.User_DateUpg = userDb.User_DateUpg.Time
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

func SelectAllUsers(page int) ([]models.UserResponse, error) {
	log.Default().Println("Start Select All users ..")
	var users []models.UserResponse

	err := DbConnect()
	if err != nil {
		log.Default().Println("Error connecting to the database..")
		return users, err
	}
	defer Db.Close()

	offset := (page * 10) - 10

	sentence := "SELECT * FROM users"
	countSentence := "SELECT count(*) FROM users"

	if page > 1 {
		sentence += " OFFSET " + strconv.Itoa(offset)
	}

	var result *sql.Rows
	result, err = Db.Query(countSentence)
	result.Next()

	var records sql.NullInt32
	err = result.Scan(&records)
	if err != nil {
		log.Default().Println("Error retrieving count of users from the database: " + err.Error())
		return users, err
	}

	log.Default().Println(sentence) //Only uncomment for debug purposes

	result, err = Db.Query(sentence)
	if err != nil {
		log.Default().Println("Error executing sentence in the database.." + err.Error())
		return users, err
	}

	for result.Next() {
		var dbUser models.User
		user := models.UserResponse{}
		err = result.Scan(
			&dbUser.UserUUID, &dbUser.User_Email, &dbUser.User_FirstName,
			&dbUser.User_LastName, &dbUser.User_Status, &dbUser.User_DateAdd,
			&dbUser.User_DateUpg,
		)
		if err != nil {
			log.Default().Println("Error parsing users: " + err.Error())
			return []models.UserResponse{}, err
		}

		user.ParseUser(dbUser)
		users = append(users, user)
	}

	log.Default().Println("Retrieve all users sucessfull..")
	return users, nil
}
