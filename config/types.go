package config

type Config struct {
	aws AwsConfig
}

type AwsConfig struct {
	Account string
	Regions []string
}
