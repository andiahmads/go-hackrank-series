package twosum_test

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	var nums = []int{3, 2, 4, 9}
	target := 11

	res := TwoSum(nums, target)

	fmt.Println(res)

}

// func TwoSum(sum []int, target int) []int {
// 	m := make(map[int]int)

// 	for i := 0; i < len(sum); i++ {

// 		j, ok := m[target-sum[i]]

// 		if ok {
// 			result := []int{j, i}
// 			return result
// 		}
// 		m[sum[i]] = i

// 	}

// 	return nil

// }

func TwoSum(sum []int, target int) []int {
	for index, left := range sum {
		// fmt.Println(index, left)
		fmt.Printf("ini index:%v, ini left:%v\n", index, left)

		for j, right := range sum {
			fmt.Printf("ini j:%v, ini right:%v\n", j, right)

			if left+right == target && index != j {
				return []int{index, j}
			}
		}
	}

	return nil
}
