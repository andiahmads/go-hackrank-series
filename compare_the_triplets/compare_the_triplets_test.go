package compare_the_triplets

import (
	"fmt"
	"testing"
)

func CompareTheTriplets(a, b [3]int) (int, int) {

	var cok, kocok int
	for i := 0; i < len(a); i++ {
		if a[i] > b[i] {
			cok += 1
		} else if a[i] < b[i] {
			kocok += 1
		}
	}

	return cok, kocok

}

func TestCompareTheTriplets(t *testing.T) {

	var andi = [3]int{3, 5, 3}
	var ahmad = [3]int{1, 4, 5}

	a, b := CompareTheTriplets(andi, ahmad)

	fmt.Println("andi = ", a, "ahmad", b)

	// assert.Equal(t, a, b, result)
}
