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
	LocalPort string
}

func InitConfig() {
	flagMap := extractFlags()

	omnivoreToken , err := getOmnivoreToken(flagMap)
	localPort := getPort(flagMap)

	if err != nil {
		log.Fatal(err)
	}

	c := Config {
		OmnivoreAuthToken: omnivoreToken,
		LocalPort: localPort,
	}

	Cfg = c
}

func extractFlags() map[string]string {
	flagMap := map[string] string {}

	var omnivoreToken string
	var secretFilePath string
	var port string

	flag.StringVar(&omnivoreToken, "t", "" , "the Omnivore API token")
	flag.StringVar(&secretFilePath, "tf", "" , "the path to the file with the Omnivore API Token")
	flag.StringVar(&port, "p", "" , "the port where the service is going to listen")
	flag.Parse()

	flagMap["t"] = omnivoreToken
	flagMap["tf"] = secretFilePath
	flagMap["p"] = port

	return flagMap
}

func getPort(flagMap map[string] string) string {
	port := flagMap["p"]

	if port == "" {
		port = os.Getenv("PORT")
		if port == "" {
			port = "8090"
		}
	}

	return port
}

func getOmnivoreToken(flagMap map[string]string) (string, error) {

	omnivoreTokenFlag := flagMap["t"]

	omnivoreToken := os.Getenv("OMNIVORE_AUTH_TOKEN")
	if omnivoreToken == "" {
		omnivoreToken = omnivoreTokenFlag
	}

	secretFilePathFlag := flagMap["tf"]
	secretFilePath := os.Getenv("OMNIVORE_AUTH_TOKEN_FILEPATH")
	if secretFilePath == "" {
		secretFilePath = secretFilePathFlag
	}

	if omnivoreToken != "" {
		log.Println("Reading secret from env var or flag")
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