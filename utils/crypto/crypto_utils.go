package crypto

import (
	"crypto/md5"
	"encoding/hex"
)

// GetMD5 converso md5 senha
func GetMD5(input string) string {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))
}
