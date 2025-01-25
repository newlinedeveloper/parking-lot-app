package services

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

// GenerateHash generates a unique hash for a string input
func GenerateHash(input string) string {
	hash := sha256.Sum256([]byte(input))
	return hex.EncodeToString(hash[:])
}

// ToUpperCase converts a string to uppercase
func ToUpperCase(input string) string {
	return strings.ToUpper(input)
}
