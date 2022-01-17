package config

type Config struct {
	_ AwsConfig
}

type AwsConfig struct {
	Account string
	Regions []string
}
