package minmaxarray

import (
	"fmt"
	"testing"
)

func TestMinMax(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}

	var sum int
	for i := 0; i < len(arr); i++ {

		sum += arr[i]
	}

	var min, max int
	for i := 0; i < len(arr); i++ {
		fmt.Println(min)
		if sum-arr[i] <= min || max == 0 {
			min = sum - arr[i]
		}
		if sum-arr[i] >= max || max == 0 {
			max = sum - arr[i]
		}

	}

	fmt.Println(min, max)

}
