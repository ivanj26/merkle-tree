package hasher

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashSHA256(data any) string {
	sha256hasher := sha256.New()

	sha256hasher.Write([]byte(data.(string)))

	hashInBytes := sha256hasher.Sum(nil)
	hashString := hex.EncodeToString(hashInBytes)

	return hashString
}
