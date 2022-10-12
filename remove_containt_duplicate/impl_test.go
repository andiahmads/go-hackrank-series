package removecontaintduplicate

import (
	"fmt"
	"testing"
)

func TestRemoveContaintDuplicateArray(t *testing.T) {
	array := []int{1, 1, 3, 3, 4, 5}

	res := RemoveContaintDuplicate(array)
	fmt.Println(res)
}

func RemoveContaintDuplicate(arr []int) []int {

	mapping := make(map[int]bool)

	res := []int{}
	for _, entry := range arr {
		if _, exist := mapping[entry]; !exist {
			fmt.Println(exist)
			mapping[entry] = true

			res = append(res, entry)
		}

	}

	return res

}
