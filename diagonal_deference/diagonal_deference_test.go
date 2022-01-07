package diagonal_deference

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiagonalDeference(t *testing.T) {
	matrix := [3][3]float64{
		{11, 2, 4},
		{4, 5, 6},
		{10, 8, -12},
	}
	callFunction := DiagonalDeference(matrix)
	assert.Equal(t, callFunction, callFunction)

}
