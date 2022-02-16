package sumsubarray_test

import (
	"fmt"
	"testing"
)

// Input   : arr[] = {1, 2, 3}
// Output  : 20
// Explanation : {1} + {2} + {3} + {2 + 3} +
//               {1 + 2} + {1 + 2 + 3} = 20

func SumOfAllSubarrays(arr []int, n int) int {

	var result, temp int = 0, 0

	for i := 0; i < n; i++ {
		temp = 0

		for j := i; j < n; j++ {
			temp += arr[j]

			fmt.Println(arr[j])

			result += temp
		}
	}

	return result

}

func TestResult(t *testing.T) {
	arry := []int{1, 2, 3}
	n := len(arry)

	res := SumOfAllSubarrays(arry, n)

	fmt.Println(res)
}
