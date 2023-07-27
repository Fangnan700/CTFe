package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5Encrypt(str string) string {
	hash := md5.Sum([]byte(str))
	return hex.EncodeToString(hash[:])
}
