package matrixtranspose_test

import (
	"fmt"
	"testing"
)

func TestMatrixTranspose(t *testing.T) {

	mtrx := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	res := transpose(mtrx)
	fmt.Println(res)

}

func transpose(matrix [][]int) [][]int {
	newARR := make([][]int, len(matrix))

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			newARR[j] = append(newARR[j], matrix[i][j])
		}
	}
	return newARR
}
