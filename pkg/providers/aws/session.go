package aws

import (
	"context"

	"github.com/spf13/viper"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

func New() (aws.Config, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(viper.GetString("aws.regionlin")))

	return cfg, err
}
