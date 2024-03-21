package encryption

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
)

func SHA256Hashing(input string) string {
	plainText := []byte(input)
	sha256Hash := sha256.Sum256(plainText)
	return hex.EncodeToString(sha256Hash[:])
}

func CompareHashedPassword(hashedPass string, plainPass string) bool {
	decodedPass, err := hex.DecodeString(hashedPass)
	if err != nil {
		log.Println("Error in decoding hashed password : " + err.Error())
		return false
	}
	plainText := []byte(plainPass)
	sha256Hash := sha256.Sum256(plainText)
	if len(decodedPass) != len(sha256Hash) {
		return false
	}

	for i := range len(decodedPass) {
		if decodedPass[i] != sha256Hash[i] {
			return false
		}
	}
	return true
}
