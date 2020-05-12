package util

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// Md5 encryption return lowercase
func Md5(str, salt string) string {
	h := md5.New()
	h.Write([]byte(str + salt))
	// h.Write([]byte(salt))
	// h.Write([]byte(salt))
	return strings.ToLower(hex.EncodeToString(h.Sum(nil)))
}
