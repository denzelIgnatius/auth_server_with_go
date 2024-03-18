package encryption

import (
	"crypto/sha256"
	"encoding/hex"
)

func SHA256Hashing(input string) string {
	plainText := []byte(input)
	sha256Hash := sha256.Sum256(plainText)
	return hex.EncodeToString(sha256Hash[:])
}
