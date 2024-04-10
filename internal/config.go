package internal

import (
	"errors"
	"log"
	"os"
	"strings"
)

var Cfg Config

type Config struct {
	OmnivoreAuthToken string
}

func InitConfig() {
	omnivoreToken , err := getOmnivoreToken()
	if err != nil {
		log.Fatal(err)
	}

	c := Config {
		OmnivoreAuthToken: omnivoreToken,
	}

	Cfg = c
}

func getOmnivoreToken() (string, error) {
	secretEnv := os.Getenv("OMNIVORE_AUTH_TOKEN")
	secretFilePath := os.Getenv("OMNIVORE_AUTH_TOKEN_FILEPATH")

	if secretFilePath != "" {
		log.Println("Reading secret from file")
		dat, err := os.ReadFile(secretFilePath)
		if err != nil {
			log.Println(err)
			return "", errors.New("unable to read Omnivore secret file")
		}

		token := string(dat)
		// Clean up secret string to avoid errors
		token = strings.Replace(token, " ", "", -1)
		token = strings.Replace(token, "\t", "", -1)
		token = strings.Replace(token, "\n", "", -1)

		return token, nil
	} else if secretEnv != "" {
		log.Println("Reading secret from env var")
		return secretEnv, nil
	}
	
	return "", errors.New("omnivore authentication is needed for this service to work, please set environment variable or secret file")
}