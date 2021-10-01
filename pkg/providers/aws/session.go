package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

func New(region string) (aws.Config, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))

	return cfg, err
}
