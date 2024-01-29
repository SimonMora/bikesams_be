package routes

import (
	"encoding/json"
	"log"

	"github.com/SimonMora/bikesams_be/database"
	"github.com/SimonMora/bikesams_be/models"
)

func UpdateUser(body string, User string) (int, string) {
	var userReq models.UserRequest
	var user models.UserResponse

	err := json.Unmarshal([]byte(body), &userReq)
	if err != nil {
		log.Default().Println("Error parsing the body request: " + err.Error())
		return 400, "Body is not User entity parseable"
	}

	if len(userReq.User_FirstName) == 0 && len(userReq.User_LastName) == 0 {
		log.Default().Println("Not enough params in the request..")
		return 400, "User name or User lastname are required to update the user"
	}

	_, exist := database.SearchUser(User)
	if !exist {
		log.Default().Println("User not found: " + User)
		return 400, "User UUID is not registered to any user in the database"
	}

	user, err = database.UpdateUser(userReq, User)
	if err != nil {
		log.Default().Println("Error saving the user update: " + err.Error())
		return 500, "Error when tried to update user in database: " + err.Error()
	}

	uBytes, errM := json.Marshal(user)
	if errM != nil {
		log.Default().Println("Error parsing the user to string response: " + errM.Error())
		return 500, "Body is not User entity parseable"
	}

	return 200, string(uBytes)
}
