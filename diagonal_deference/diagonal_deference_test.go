package diagonal_deference

import (
	"fmt"
	"testing"
)

func TestDiagonalDeference(t *testing.T) {
	matrix := [3][3]float64{
		{11, 2, 4},
		{4, 5, 6},
		{10, 8, -12},
	}
	callFunction := DiagonalDeference(matrix)
	fmt.Println(callFunction)

}
