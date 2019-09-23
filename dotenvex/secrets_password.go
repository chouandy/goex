package dotenvex

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

var secretsPassword = os.Getenv("SECRETS_PASSWORD")

var (
	secretsPasswordPsName   = os.Getenv("SECRETS_PASSWORD_PS_NAME")
	secretsPasswordPsRegion = os.Getenv("SECRETS_PASSWORD_PS_REGION")
)

// GetSecretsPassword get secrects password
func GetSecretsPassword() string {
	// Password from env
	if len(secretsPassword) > 0 {
		return secretsPassword
	}

	// Check parameter store envs
	if len(secretsPasswordPsName) > 0 && len(secretsPasswordPsRegion) > 0 {
		// Get password from aws parameter store
		password, err := GetSecretsPasswordFromParameterStore()
		if err != nil {
			fmt.Printf("[GetSecretsPassword] [GetSecretsPasswordFromParameterStore] %s\n", err)
			return ""
		}

		return password
	}

	return ""
}

// GetSecretsPasswordFromParameterStore get secrets password fromparameterstore
func GetSecretsPasswordFromParameterStore() (string, error) {
	// new aws ssm client
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return "", err
	}
	cfg.Region = secretsPasswordPsRegion
	client := ssm.New(cfg)

	// New input
	input := &ssm.GetParameterInput{
		Name:           aws.String(secretsPasswordPsName),
		WithDecryption: aws.Bool(true),
	}

	// New request
	request := client.GetParameterRequest(input)

	// Send request
	resp, err := request.Send(context.Background())
	if err != nil {
		return "", err
	}

	return aws.StringValue(resp.Parameter.Value), nil
}
