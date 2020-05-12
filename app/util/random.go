package util

import (
	"gin-user-center/app/common"
	"math/rand"
	"time"
)

// Generates a random string of length
func RandString(length int) string {
	// 设置随机种子
	rand.Seed(time.Now().UnixNano())
	bytes := make([]byte, length)
	for i := 0; i < length; i++ {
		b := common.BASE_SALT[rand.Intn(len(common.BASE_SALT))]
		bytes[i] = byte(b)
	}
	return string(bytes)
}
