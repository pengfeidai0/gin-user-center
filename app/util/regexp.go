package util

import "regexp"

// checkk mobile format
func CheckPhone(phone string) bool {
	reg := `^1(3|4|5|6|7|8)\d{9}$`
	rgx := regexp.MustCompile(reg)
	return rgx.MatchString(phone)
}
