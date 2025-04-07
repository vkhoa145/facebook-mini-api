package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math/big"
)

func GenerateRandomString(length int) (string, error) {
	bytes := make([]byte, length/2)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	return hex.EncodeToString(bytes), nil
}

func GenerateVerifyCode() string {
	num, _ := rand.Int(rand.Reader, big.NewInt(1000000))
	return fmt.Sprintf("%06d", num.Int64())
}
