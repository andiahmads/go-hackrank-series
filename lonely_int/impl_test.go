package lonelyint_test

import (
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {
	arr := []int{1, 2, 3, 4, 3, 2, 1}

	res := LoneylyInt(arr)
	fmt.Println(res)
}

func LoneylyInt(x []int) int {

	intMap := make(map[int]int)

	for _, val := range x {
		intMap[val]++
	}

	for _, val := range x {
		if intMap[val] == 1 {
			return val
		}
	}

	return 0
}
