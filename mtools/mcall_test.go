package mtools

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMGet(t *testing.T) {
	//inputs := []string{"a", "b", "c"}
	//outputs := make([]string, 3)
	inputs := []int{1, 2, 3}
	outputs := make([]int, len(inputs))
	err := MGet(inputs, &outputs, func(i interface{}) interface{} {
		return i.(int) + 10
	})
	assert.Equal(t, 11, outputs[0])

	input2 := []string{"a", "b", "c"}
	output2 := make([]string, 3)
	err = MGet(input2, &output2, func(i interface{}) interface{} {
		return i.(string) + "1"
	})
	assert.Nil(t, err)
	assert.Equal(t, "a1", output2[0])

	type A struct {
		A string
	}
	input3 := []string{"a", "b", "c"}
	output3 := make([]A, 3)
	err = MGet(input3, &output3, func(i interface{}) interface{} {
		return A{A: i.(string)}
	})
	assert.Nil(t, err)
	assert.Equal(t, "a", output3[0].A)

	input4 := []string{"a", "b", "c"}
	var output4 []A
	err = MGet(input4, &output4, func(i interface{}) interface{} {
		return A{A: i.(string)}
	})
	assert.Error(t, OutputsLenErr, err)
}

func BenchmarkMGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input2 := []string{"a", "b", "c"}
		output2 := make([]string, 3)
		_ = MGet(input2, &output2, func(i interface{}) interface{} {
			return i.(string) + "1"
		})
	}
}

func BenchmarkMCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := []string{"a", "b", "c"}
		_ = MCall(input, func(i interface{}) {
			// nothing
		})
	}
}
