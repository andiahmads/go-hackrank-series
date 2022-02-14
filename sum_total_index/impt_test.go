package sumtotalindex

import (
	"fmt"
	"sort"
	"testing"
)

func SumTotal(list []int) map[int]int {

	//sort array asc

	var countSameIndex = make(map[int]int)
	for _, v := range list {

		_, exist := countSameIndex[v]
		if exist {
			countSameIndex[v] += 1
		} else {
			countSameIndex[v] = 1
		}
	}

	return countSameIndex

}

func TestTestingSum(t *testing.T) {
	var arr = []int{700, 400, 200, 800, 700, 100, 200, 100, 300}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})

	count := SumTotal(arr)

	for key, val := range count {
		fmt.Println(key, "berjumlah = ", val)
	}
}
