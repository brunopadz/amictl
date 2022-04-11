package config

type Config struct {
	Aws AwsConfig
}

type AwsConfig struct {
	Account string
	Regions []string
}
