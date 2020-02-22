package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type ApiConfig struct {
	Api        string
	ApiPrefix  string
	CustomerId string
	PrivateKey string
}

const apiConfigFile = "./etc/config.json"

func GetApiConfig() ApiConfig {
	apiFile, _ := os.Open(apiConfigFile)
	defer apiFile.Close()
	apiConfiguration := ApiConfig{}
	apiErr := json.NewDecoder(apiFile).Decode(&apiConfiguration)

	if apiErr != nil {
		fmt.Println("error:", apiErr)
	}

	return apiConfiguration
}
