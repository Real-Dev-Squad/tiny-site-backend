package utils

import (
	"crypto/md5"
	"encoding/hex"
	"time"
)

func GenerateMD5Hash(url string) string {
	url = url + time.Nanosecond.String()
	hash := md5.New()
	hash.Write([]byte(url))
	return hex.EncodeToString(hash.Sum(nil))[:8]
}
