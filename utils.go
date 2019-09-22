package utils

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/google/uuid"
	"github.com/lithammer/shortuuid"
)

// Sha256 to generate sha256
func Sha256(data string) string {
	h := sha256.New()
	h.Write([]byte(data))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

// GenerateUUID - Function to generate UUID
func GenerateUUID() string {
	return uuid.New().String()
}

// GenerateShortID - Function to Generate Short UUID
func GenerateShortID() string {
	return shortuuid.New()
}
