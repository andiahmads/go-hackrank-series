package maximumsubarraysizek_test

import (
	"fmt"
	"math"
	"testing"
)

func MaximumSubArrayK(arr []int, k int) int {

	var maxSum, windowsSum int

	for i := 0; i <= len(arr)-k; i++ {
		windowsSum = 0

		lenarr := len(arr) - k

		fmt.Println(lenarr)

		t := i + k
		fmt.Println(t)
		for j := i; j < t; j++ {

			windowsSum += arr[j] // 2, 1, 5, 1, 3, 2
		}

		/*
			step 1

			j = 0 = 2
			j = 1 = 1
			j = 2= 5
		*/

		/*
			step 2

			j = 1 = 1
			j = 2 = 5
			j = 3= 1
		*/

		maxSum = int(math.Max(float64(maxSum), float64(windowsSum)))

	}

	return maxSum

}

func TestMaximumSubArrayK(t *testing.T) {

	arr := []int{2, 1, 5, 1, 3, 2}

	k := 4

	res := MaximumSubArrayK(arr, k)

	fmt.Println(res)

}
