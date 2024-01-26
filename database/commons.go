package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/SimonMora/bikesams_be/models"
	"github.com/SimonMora/bikesams_be/secrets"
	_ "github.com/go-sql-driver/mysql"
)

var SecretModel models.SecretRdsJson
var err error
var Db *sql.DB

func ReadSecrets() error {
	SecretModel, err = secrets.GetSecret(os.Getenv("SecretName"))
	return err
}

func DbConnect() error {
	log.Default().Println("Start connection to db process..")
	if len(SecretModel.Username) == 0 {
		err = ReadSecrets()
		if err != nil {
			log.Default().Println("Error retrieving database secrets from SecretsManager")
			return err
		}
	}

	Db, err = sql.Open("mysql", ConnStr(SecretModel))
	if err != nil {
		log.Default().Println(err.Error())
		return err
	}

	err = Db.Ping()
	if err != nil {
		log.Default().Println(err.Error())
		return err
	}

	log.Default().Println("Successfully connected to db..")
	return nil
}

func ConnStr(credentials models.SecretRdsJson) string {
	var hostName, dbName, dbUser, password string
	dbUser = credentials.Username
	password = credentials.Password
	hostName = credentials.Host
	dbName = "bikesams"
	return fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?allowCleartextPasswords=true&parseTime=true",
		dbUser, password, hostName, dbName,
	)
}

func IsUserAdminValidate(userUUID string) (bool, string) {
	log.Default().Println("Start user status validation..")

	err := DbConnect()
	if err != nil {
		log.Default().Println("Error connecting to the database..")
		return false, err.Error()
	}
	defer Db.Close()

	sentence := "SELECT 1 FROM users where User_UUID = '" + userUUID + "' AND User_Status = 0"
	//log.Default().Println(sentence) //Only uncomment for debug purposes

	rows, errSql := Db.Query(sentence)
	if errSql != nil {
		log.Default().Println("Error executing SQL sentence..")
		return false, errSql.Error()
	}

	var queryResult string
	rows.Next()
	rows.Scan(&queryResult)

	log.Default().Println("User validation return status: " + queryResult)

	if queryResult != "1" {
		return false, "User is not an admin."
	}

	return true, ""
}
