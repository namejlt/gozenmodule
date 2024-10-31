package mtools

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxHans(t *testing.T) {
	testCases := []struct {
		giveStr  string
		giveMax  int
		wantStr  string
		wantBool bool
	}{
		{
			giveStr:  "",
			giveMax:  1,
			wantStr:  "",
			wantBool: false,
		},
	}

	for _, test := range testCases {
		output, ok := MaxHans(test.giveStr, test.giveMax)
		assert.Equal(t, test.wantBool, ok)
		assert.Equal(t, test.wantStr, output)
	}
}
