package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// The Config for the program
type Config struct {
	DBAddress    string `json:"DBAddress"`
	DBUsername   string `json:"DBUsername"`
	DBPassword   string `json:"DBPassword"`
	DBName       string `json:"DBName"`
	DBParameters string `json:"DBParameters"`
}

// LoadConfig file
func LoadConfig(path string) (Config, error) {
	var config Config

	data, err := ioutil.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			data, err = json.Marshal(&config)
			if err != nil {
				return config, err
			}
			err = ioutil.WriteFile(path, data, 0660)
			if err != nil {
				return config, err
			}
			log.Println("A blank config file has been created, please fill it in.")
		}
		return config, err
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}
	return config, nil
}

// LoadEnvironmentConfig gets config values from the environment
func LoadEnvironmentConfig() Config {
	Environment := make(map[string]string)
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		Environment[pair[0]] = strings.Join(pair[1:], "=")
	}
	return Config{
		DBAddress:    Environment["DBAddress"],
		DBUsername:   Environment["DBUsername"],
		DBPassword:   Environment["DBPassword"],
		DBName:       Environment["DBName"],
		DBParameters: Environment["DBParameters"],
	}
}
