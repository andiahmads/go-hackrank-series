package diagonal_deference

import (
	"fmt"
	"math"
	"testing"
)

func Test2DiagonalDeference(t *testing.T) {
	matrix := [][]int{
		{11, 2, 4},
		{4, 5, 6},
		{10, 8, -12},
	}

	res := diagonal_deference2(matrix)

	fmt.Println(res)

}

func diagonal_deference2(arr [][]int) int {

	res1 := 0
	res2 := 0

	for i := 0; i < len(arr); i++ {

		fmt.Println(arr[i])

		res1 += arr[i][i]

		/*
			loop sepanjang len arr == 3

			loop 1
			 i baris= 0  colom 0 = 11
			 i baris = 1 colom 1 = 5
			 i baris = 2 colom 2 = -12
		*/
		res2 += arr[i][len(arr)-i-1]

	}

	result := math.Abs(float64(res1) - float64(res2))

	return int(result)
}
