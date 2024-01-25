package auth

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"strings"
	"time"
)

type TokenJson struct {
	Sub       string
	Event_Id  string
	Token_Use string
	Scope     string
	Auth_Time int
	Iss       string
	Exp       int
	Iat       int
	Client_id string
	Username  string
}

func TokenValidate(token string) (bool, error, string) {
	log.Default().Println("Start token valdiation process.. Retrieving token properties.")
	splittedToken := strings.Split(token, ".")

	if len(splittedToken) != 3 {
		log.Default().Println("Invalid Token.")
		return false, nil, "Invalid Token."
	}

	userInfo, err := base64.StdEncoding.WithPadding(base64.NoPadding).DecodeString(splittedToken[1])
	if err != nil {
		log.Default().Println("Error decoding the user information: " + err.Error())
		return false, err, err.Error()
	}

	var tknInfo TokenJson
	err = json.Unmarshal(userInfo, &tknInfo)
	if err != nil {
		log.Default().Println("Error unmarshalling the token information: " + err.Error())
		return false, err, err.Error()
	}

	now := time.Now()
	tknExp := time.Unix(int64(tknInfo.Exp), 0)
	if tknExp.Before(now) {
		log.Default().Println("Expired Token. Expiration date: " + tknExp.String())
		return false, nil, "Expired Token."
	}

	log.Default().Println("Token validation ok, returning user id..")
	return true, nil, string(tknInfo.Sub)
}
