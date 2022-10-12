package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// func main() {
// 	var data = [6]int{1, 2, 3, 4, 10, 11}
// 	var res int
// 	res = simple_array_sum.ArraySum(data, 6)
// 	fmt.Printf("Final result is: %d ", res)
// }

// func main() {

// 	var andi = [3]int{3, 5, 3}
// 	var ahmad = [3]int{1, 4, 5}

// 	a, b := compare_the_triplets.CompareTheTriplets(andi, ahmad)
// 	fmt.Println(a, b)

// }

// func main() {
// 	// var matrix = diagonal_deference.DiagonalDeference
// 	matrix := [3][3]float64{
// 		{11, 2, 4},
// 		{4, 5, 6},
// 		{10, 8, -12},
// 	}
// 	callFunction := diagonal_deference.DiagonalDeference(matrix)
// 	fmt.Println(callFunction)

// }

func main() {
	var line string
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line = scanner.Text()

		world := strings.Fields(line)

		for i, v := range world {
			str := v[i:]
			str2 := v[i:]
			fmt.Println(string(str))
			fmt.Println(string(str2))

			// toUpper := strings.ToUpper(str)
		}

	}

}
