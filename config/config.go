package config

import (
	"context"
	"os"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	
)


type Config struct {
	AWSConfig *aws.Config
}

func NewConfig() *Config{
	ctx := context.Background()
	log.Println(os.Getenv("AWS_PROFILE"))
	awsConfig,err := config.LoadDefaultConfig(
		ctx,config.WithDefaultRegion("ap-northeast-1"),
		config.WithSharedConfigProfile(os.Getenv("AWS_PROFILE")),
		config.WithAssumeRoleCredentialOptions(func(options *stscreds.AssumeRoleOptions){
			options.TokenProvider = func() (string,error) {
				return os.Getenv("MFA_TOKEN"),nil			}
		}	),
	)
  if err != nil {
		panic(err)	
	}	

	return &Config{
		AWSConfig: &awsConfig,
	} 
	
}
