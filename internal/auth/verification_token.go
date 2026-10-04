package auth

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateVerificationToken() (string, error) {

	b := make([]byte, 32)

	_, err := rand.Read(b)

	if err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(b), nil
}
