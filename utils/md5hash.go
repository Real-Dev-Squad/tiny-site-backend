package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"time"
)

func GenerateMD5Hash(url string) string {
	currentTime := time.Now().UTC()
	url += strconv.FormatInt(currentTime.UnixNano(), 10) 
	hash := md5.New()
	hash.Write([]byte(url))
	return hex.EncodeToString(hash.Sum(nil))[:5]
}
