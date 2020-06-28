package main

import (
	"config"
	"fmt"
	"os"
	"router"
)

func getConfig(mode string) *config.Config {
	if mode == "development" {
		return &config.Config{
			Type: "development",
			S3Config: config.S3Config{
				Endpoint: "sfo2.digitaloceanspaces.com",
				AccessID: "B7DGWCR5E3MOXTQANMRX",
				Secret:   "pKnaWLIcGvNiRe5hCa9y8dCx6M/dXap96w8fd4nT/1k",
				UseSSL:   true,
			},
			Tokens: config.Tokens{
				FileKey: "./dev_secrets/file_key.txt",
			},
			Host: "Domain",
			Port: ":2000",
		}
	}

	return &config.Config{
		Type: "production",
		S3Config: config.S3Config{
			Endpoint: "sfo2.digitaloceanspaces.com",
			AccessID: "B7DGWCR5E3MOXTQANMRX",
			Secret:   "pKnaWLIcGvNiRe5hCa9y8dCx6M/dXap96w8fd4nT/1k",
			UseSSL:   true,
		},
		Tokens: config.Tokens{
			FileKey: "/run/secrets/file_key",
		},
		Host: "Domain",
		Port: ":2000",
	}
}

func main() {

	//Default Dev config
	configType := "development"

	//Check if production flag is passed.
	if len(os.Args) > 1 {
		if os.Args[1] == "-production" {
			configType = "production"
		}
	}

	//Get correct runtime config
	config := getConfig(configType)

	//Start router
	err := router.Router{}.Init(config)
	if err != nil {
		fmt.Println(err)
		return
	}

}
