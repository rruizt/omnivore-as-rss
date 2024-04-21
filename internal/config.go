package internal

import (
	"errors"
	"log"
	"os"
	"flag"
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
	var omnivoreTokenFlag string
	var secretFilePathFlag string

	flag.StringVar(&omnivoreTokenFlag, "t", "" , "the Omnivore API token")
	flag.StringVar(&secretFilePathFlag, "tf", "" , "the path to the file with the Omnivore API Token")

	flag.Parse()

	omnivoreToken := os.Getenv("OMNIVORE_AUTH_TOKEN")
	if omnivoreToken == "" {
		omnivoreToken = omnivoreTokenFlag
	}

	secretFilePath := os.Getenv("OMNIVORE_AUTH_TOKEN_FILEPATH")
	if secretFilePath == "" {
		secretFilePath = secretFilePathFlag
	}

	if omnivoreToken != "" {
		log.Println("Reading secret from env var")
		return omnivoreToken, nil
	} else if secretFilePath != "" {
		log.Println("Reading secret from file")

		dat, err := os.ReadFile(secretFilePath)
		if err != nil {
			log.Println(err)
			return "", errors.New("unable to get Omnivore API Key: filepath was provided but couldn't read it")
		}

		token := string(dat)
		// Clean up secret string to avoid errors
		token = strings.Replace(token, " ", "", -1)
		token = strings.Replace(token, "\t", "", -1)
		token = strings.Replace(token, "\n", "", -1)

		return token, nil
	}

	return "", errors.New("omnivore API key is needed for this application to work, please set environment variable or flag")
}