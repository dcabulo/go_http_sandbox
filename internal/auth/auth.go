package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKEY extrack an API Key from
// the heders of an HTTP request
// Example:
// Authorization: ApiKey {APIKEY}

func GetAPIKey(headers http.Header) (string, error) {
	apiKey := headers.Get("Authorization")
	if apiKey == "" {
		return "", errors.New("no authentication info around")
	}

	vals := strings.Split(apiKey, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed auth header")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first part of the auth header")
	}
	return vals[1], nil
}
