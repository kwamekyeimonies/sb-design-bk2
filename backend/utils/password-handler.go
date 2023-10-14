package utils

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func HashPasswordHandler(password, key string) (string, error) {
	data := []byte(password + key)
	hash := sha256.Sum256(data)
	encodedHash := base64.StdEncoding.EncodeToString(hash[:])

	return encodedHash, nil
}

func CheckPasswordHashHandler(password, hash, key string) error {
	newHash, err := HashPasswordHandler(password, key)
	if err != nil {
		return err
	}
	if newHash != hash {
		return fmt.Errorf("invalid password")
	}

	return nil
}
