package sun

import (
	"fmt"
	"testing"
)

func TestSun(t *testing.T) {
	building := []int{2, 6, 7, 99, 53, 4, 1}
	res := Sun(building)
	fmt.Println(res)
}

func Sun(arr []int) int {

	max := 0
	index := 0
	var sisa []int

	for i, v := range arr {
		if v > max {
			max = v

			index = i
		}

		sisa = arr[index+1:]
	}

	return len(sisa)

}
