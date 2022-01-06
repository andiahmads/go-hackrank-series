package simple_array_sum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArraySum(t *testing.T) {
	var data = [6]int{1, 2, 3, 4, 10, 116}
	var res int;

	res = ArraySum(data, 6)
	assert.Equal(t, 136, res)
}
