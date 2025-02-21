package auth

import (
	"errors"
	"net/http"
	"strings"
	"testing"
)

var ErrNoAuthHeaderIncluded = errors.New("no authorization header included")

// GetAPIKey -
func GetAPIKey(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", ErrNoAuthHeaderIncluded
	}
	splitAuth := strings.Split(authHeader, " ")
	if len(splitAuth) < 2 || splitAuth[0] != "ApiKey" {
		return "", errors.New("malformed authorization header")
	}

	return splitAuth[1], nil
}

func TestGetAPIKey(t *testing.T) {
	req, err := http.NewRequest("GET", "http://1.1.1.1/", nil)
	if err != nil {
		t.Fatalf("Error creating the request")
	}

	_, err = GetAPIKey(req.Header)
	if err != ErrNoAuthHeaderIncluded {
		t.Fatalf("Expected ErrNoAuthHeaderIncluded on request with empty header")
	}
}
