package sdk

type Config struct {
	AccessKey string
	SecretKey string
}

func NewConfig(accessKey, secretKey string) *Config {
	return &Config{
		AccessKey: accessKey,
		SecretKey: secretKey,
	}
}
