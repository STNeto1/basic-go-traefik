package utils

import (
	"errors"
	"strings"
)

func GetToken(tokenStr string) (string, error) {
	authHeaderParts := strings.Fields(tokenStr)
	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
		return "", errors.New("not valid")
	}

	return authHeaderParts[1], nil
}
