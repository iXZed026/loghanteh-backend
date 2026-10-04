package security

import (
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	appErrors "loghanteh-project/internal/errors"

	"golang.org/x/crypto/bcrypt"
)

const bcryptCost = bcrypt.DefaultCost

func HashPassword(password string) (string, error) {
	if password == "" {
		return "", appErrors.ErrEmptyValue
	}

	hashed, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcryptCost,
	)

	if err != nil {
		return "", err
	}

	return string(hashed), nil
}

func CheckPasswordHash(hashedPassword, password string) bool {
	if hashedPassword == "" || password == "" {
		return false
	}

	err := bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword),
		[]byte(password),
	)

	return err == nil
}

func HashToken(token string) string {
	sum := sha256.Sum256([]byte(token))
	return hex.EncodeToString(sum[:])
}

func CompareTokenHash(hashedToken, rawToken string) bool {
	if hashedToken == "" || rawToken == "" {
		return false
	}

	computedHash := HashToken(rawToken)

	return subtle.ConstantTimeCompare(
		[]byte(hashedToken),
		[]byte(computedHash),
	) == 1
}
