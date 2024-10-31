package mtools

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"math/rand"
)

func Md5(input []byte) string {
	hash := md5.New()
	hash.Write(input)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

// MaxHans 汉字最大长度截取
func MaxHans(input string, max int) (output string, ok bool) {
	i := []rune(input)
	if len(i) > max {
		return string(i[:max]), false
	}
	return input, true
}

func Base64Encode(input string) string {
	return base64.StdEncoding.EncodeToString([]byte(input))
}

func UtilCryptoMd5(s string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(s))
	cipherStr := md5Ctx.Sum(nil)
	return base64.StdEncoding.EncodeToString(cipherStr)
}

func RandomStr(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyz0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}
