package compare_the_triplets

import (
	"fmt"
	"testing"
)

func CompareArrayScore(a, b []int) (int, int) {

	var arr1 int
	var arr2 int

	for i := 0; i < len(a); i++ {
		if a[i] > b[i] {
			arr1 += 1
		} else if b[i] > a[i] {
			arr2 += 1
		}
	}

	return arr1, arr2

}

func TestCompareArrayScores(t *testing.T) {
	arr1 := []int{1, 5, 7, 1, 5, 7}
	arr2 := []int{2, 2, 3, 1, 5, 7}

	res, res2 := CompareArrayScore(arr1, arr2)

	fmt.Println("arr1 = ", res, "arr2 = ", res2)
}
