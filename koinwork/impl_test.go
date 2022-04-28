package koinwork

import (
	"fmt"
	"testing"
)

func Solution(a []int) int {

	var minval = len(a)
	for i, v := range a {
		if v%2 == 0 {
			if i == 0 {
				minval = v
			}
			if i < minval {
				minval = v

			}

		}

	}
	return minval
}

func TestKoinWork(t *testing.T) {

	arr := []int{1, 3, 2, 6, 4, 1, 2}

	res := Solution(arr)
	fmt.Println(res)
	s := fmt.Sprintf("%b", res)
	fmt.Println(s)

}
