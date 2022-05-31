package removecontaintduplicate

import (
	"fmt"
	"testing"
)

func TestRemoveContaintDuplicateArray(t *testing.T) {
	array := []int{1, 2, 2, 4, 5, 77, 7, 2, 1, 41, 3, 2, 5, 9, 5}

	res := RemoveContaintDuplicate(array)
	fmt.Println(res)
}

func RemoveContaintDuplicate(arr []int) []int {

	mapping := make(map[int]bool)

	res := []int{}
	for _, entry := range arr {
		if _, exist := mapping[entry]; !exist {
			mapping[entry] = true

			res = append(res, entry)
		}

	}

	return res

}
