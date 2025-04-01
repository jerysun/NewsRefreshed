package auth

import (
	"errors"
	"net/http"
	"strings"
)

var ErrNoAuthHeaderIncluded = errors.New("no authorization header included")

// GetAPIKey extracts an API Key from the headers of an HTTP request
// Example: Authorization: ApiKey {api_key}
func GetAPIKey(headers http.Header) (string, error) {
  authHeader := headers.Get("Authorization")
  if authHeader == "" {
    return "", ErrNoAuthHeaderIncluded
  }

  splitAuth := strings.Split(authHeader, " ")
  if len(splitAuth) < 2 || splitAuth[0] != "ApiKey" { // splitAuth[0] is "ApiKey", splitAuth[1] is the 64-byte hex value
    return "", errors.New("malformed authorization header")
  }

  return splitAuth[1], nil
}