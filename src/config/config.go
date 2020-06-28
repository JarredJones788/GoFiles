package config

//S3Config - struct to connect to a s3/spaces endpoint
type S3Config struct {
	Endpoint string
	AccessID string
	Secret   string
	UseSSL   bool
}

//Tokens - struct for all tokens
type Tokens struct {
	FileKey string
}

//Config - structure of config
type Config struct {
	Type     string
	S3Config S3Config
	Tokens   Tokens
	Host     string
	Port     string
}
