package print_diagonal_matrix

import "fmt"

func PrintMatrixDiagonal(matrix [3][3]int) int {
	var result int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i+j == 2 {
				fmt.Printf("%d ", matrix[i][i])
			} else {
				fmt.Printf("")
			}
		}
		fmt.Printf("\n")
	}

	return result
}
