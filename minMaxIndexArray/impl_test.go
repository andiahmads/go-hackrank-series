package minmaxindexarray

import (
	"fmt"
	"testing"
)

func minMaxIndexArray(arr []int) (int, int) {

	var min, max int
	fmt.Println(min)
	for i, val := range arr {
		switch {
		case i == 0:
			min = val
		case i > max:
			max = val
		case i < min:
			min = val
		}
	}
	return min, max
}

func TestMinMaxIndexArray(t *testing.T) {
	var arr = []int{1, 1, 2, 3, 4, 5, 21}

	min, max := minMaxIndexArray(arr)

	fmt.Println(min, max)

}
