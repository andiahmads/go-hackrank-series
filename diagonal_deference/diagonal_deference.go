package diagonal_deference

import (
	"math"
)

func DiagonalDeference(matrix [3][3]float64) float64 {

	diagonal1 := 0.0
	diagonal2 := 0.0

	for i := 0; i < len(matrix); i++ {
		diagonal1 += matrix[i][i]

		diagonal2 += matrix[i][len(matrix)-i-1]

	}
	return math.Abs(diagonal1 - diagonal2)

}
