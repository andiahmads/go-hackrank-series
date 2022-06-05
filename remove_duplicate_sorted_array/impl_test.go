package removeduplicatesortedarray_test

import (
	"fmt"
	"testing"
)

func TestRemoveContaintDuplicateArray(t *testing.T) {
	arr := []int{1,1,2,2,4}
	res := removecontaintduplicate(arr)
	fmt.Println(res)

}

func removecontaintduplicate(nums []int) int {
	length, newLength := len(nums), 1

	for i := 1; i < length; i++ {
		if nums[i-1] != nums[i] {
			nums[newLength] = nums[i]
			newLength++
		}
	}
	return newLength
}
