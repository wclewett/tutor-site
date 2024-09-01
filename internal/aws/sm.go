package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

type SMClient struct {
  client *secretsmanager.Client
}

func GetSMClient(region string) (*SMClient, error) {
  sm := &SMClient{}
  cfg, err := authorizeSDK(region)
  sm.client = secretsmanager.NewFromConfig(*cfg)
  return sm, err 
}

func (sm *SMClient) GetSecret(secretId *string) (*string, error) {
  var secret string
  input := &secretsmanager.GetSecretValueInput{SecretId: secretId, }
  output, err := sm.client.GetSecretValue(context.Background(), input)
  if err != nil { return &secret, err }
  secret = *output.SecretString
  return &secret, nil
}
