package simple_array_sum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func ArraySumTest(t *testing.T) {
	var data = [6]int{1, 2, 3, 4, 10, 11}
	var res int

	res = ArraySum(data, 6)
	assert.Equal(t, 31, res)

}
