package utils

import (
	"crypto/md5"
	"encoding/hex"
)

//Encriptar para md5
func Md5(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
