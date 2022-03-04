package deleteduplicatearray

import (
	"fmt"
	"testing"
)

func TestDeleteDuplicateArray(t *testing.T) {
	var numbers = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}
	res := DeleteDuplicateArray(numbers)
	fmt.Println(res)
}

func DeleteDuplicateArray(arr []int) []int {
	var mappings = make(map[int]bool)

	Newlist := []int{}

	for _, val := range arr {
		if _, value := mappings[val]; !value {
			mappings[val] = true
			Newlist = append(Newlist, val)
		}

	}

	return Newlist
}
