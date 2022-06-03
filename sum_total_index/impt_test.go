package sumtotalindex

import (
	"fmt"
	"sort"
	"testing"
)

func SumTotal(list []int) map[int]int {

	//maping object
	var countSameIndex = make(map[int]int)

	for _, v := range list {

		// varibel dibawah untuk cek total index yg sama
		// yg akan mengembalikan nilai true & false
		_, exist := countSameIndex[v]

		//jika elemet sama dengan true +1
		if exist {
			countSameIndex[v] += 1
		} else {
			countSameIndex[v] = 1 // jika false ttp 1
		}
	}

	return countSameIndex

}

func TestTestingSum(t *testing.T) {
	var arr = []int{700, 400, 200, 800, 700, 100, 200, 100, 300}

	//sorting array
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})

	count := SumTotal(arr)

	//cetak array
	for key, val := range count {
		fmt.Println(key, "berjumlah :%d ", val)
	}
}
