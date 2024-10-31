package mtools

import (
	"strings"
	"testing"
)

func TestStringTenSlice(t *testing.T) {
	testCases := map[string]int{
		"1":                       1,
		"1,2,3,4,5,6,7,8,9":       1,
		"1,2,3,4,5,6,7,8,9,10":    1,
		"1,2,3,4,5,6,7,8,9,10,11": 2,
		"1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20":    2,
		"1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21": 3,
	}

	for testCase, expectedLen := range testCases {
		list := strings.Split(testCase, ",")
		data := StringTenSlice(list)
		if len(data) != expectedLen {
			t.Errorf("test failed, expect %d, got %d", expectedLen, len(data))
		}
	}
}
