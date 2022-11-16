package encryption

import (
	"crypto/sha256"
)

// Convert string into SHA256 hash
func Shasum(key string) []byte {
	data := []byte(key)
	hash := sha256.New()
	hash.Write(data)
	sum := hash.Sum(nil)
	return sum
}
