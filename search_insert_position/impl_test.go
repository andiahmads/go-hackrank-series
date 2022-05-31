package searchinsertposition_test

import (
	"fmt"
	"testing"
)

func TestSearchInsertPosition(t *testing.T) {
	num := []int{1, 3, 5, 6}
	target := 2

	res := searchInsert(num, target)
	fmt.Printf("hasil:%v \n", res)
}

func searchInsert(nums []int, target int) int {
	low := 0

	for low < len(nums) {
		fmt.Println(nums[low])

		if nums[low] < target {
			low++
		} else {
			return low
		}
	}

	return low
}
