package mtools

import "unicode/utf8"

// FilterEmoji 过滤emoji表情
func FilterEmoji(input string) (output string) {
	for _, value := range input {
		_, size := utf8.DecodeRuneInString(string(value))
		if size <= 3 {
			output += string(value)
		}
	}
	return
}

// HasEmoji 判断是否有表情
func HasEmoji(input string) (hasEmoji bool) {
	for _, value := range input {
		_, size := utf8.DecodeRuneInString(string(value))
		if size > 3 {
			return true
		}
	}
	return false
}
