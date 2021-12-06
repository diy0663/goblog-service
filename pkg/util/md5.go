package util

import (
	"crypto/md5"
	"encoding/hex"
)

// 封装MD5
func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))
	return hex.EncodeToString(m.Sum(nil))
}
