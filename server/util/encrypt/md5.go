package encrypt

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5Encrypt(str string) string {
	hash := md5.Sum([]byte(str))
	md5Str := hex.EncodeToString(hash[:])

	return md5Str
}
