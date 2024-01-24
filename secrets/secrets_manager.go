package secrets

import (
	"encoding/json"
	"log"

	"github.com/SimonMora/bikesams_be/aws_go"
	"github.com/SimonMora/bikesams_be/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func GetSecret(secretName string) (models.SecretRdsJson, error) {
	var secrets models.SecretRdsJson
	log.Default().Printf("Start retrieving secrets with name: %s", secretName)

	smClient := secretsmanager.NewFromConfig(aws_go.Cfg)
	jsonSecret, err := smClient.GetSecretValue(aws_go.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	})

	if err != nil {
		log.Default().Printf("Error retrieving secrets from secret manager: %v", err.Error())
		return secrets, err
	}

	json.Unmarshal([]byte(*jsonSecret.SecretString), &secrets)
	log.Default().Printf("Secret %s retrieved successfully.", secretName)

	return secrets, nil
}
