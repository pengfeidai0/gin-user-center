package util

import (
	"gin-user-center/app/common"
	"math/rand"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
)

func init() {
	// 设置随机种子
	rand.Seed(time.Now().UnixNano())
}

// Generates a random string of length
func RandString(length int) string {
	bytes := make([]byte, length)
	for i := 0; i < length; i++ {
		b := common.BASE_SALT[rand.Intn(len(common.BASE_SALT))]
		bytes[i] = byte(b)
	}
	return string(bytes)
}

// Generates uuid v4
func Uuid() string {
	u := uuid.Must(uuid.NewV4(), nil).String()
	return strings.ReplaceAll(u, "-", "")
}
