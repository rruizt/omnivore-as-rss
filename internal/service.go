package internal

import (
	"errors"
	"log"
	"strings"
)

func getFeed() (string, error) {
	err := queryOmnivore()

	if err != nil {
		log.Println("connection error on GraphQL Query: ", err)
		return "", errors.New("connection to Omnivore error. See logs")
	}

	// Check errorCodes from GraphQL API
	errorCodes := SearchQuery.Search.SearchError.ErrorCodes
	if len(errorCodes) != 0 {
		errorCodesString := strings.Join(errorCodes[:], ",")
		return "", errors.New("Errors returned from Omnivore GraphQL API: " + errorCodesString)
	}

	return generateFeed()
}
