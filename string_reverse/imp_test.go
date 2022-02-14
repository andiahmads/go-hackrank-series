package stringreverse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func ReverseString(char string) string {
	temp := ""
	for i := 0; i < len(char); i++ {
		arr := string(char[i])
		temp = arr + temp
	}

	return temp
}

func TestStringReverse(t *testing.T) {

	char := ReverseString("abc")

	assert.Equal(t, "cba", char)

}
