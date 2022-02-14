package plusminus_test

import (
	"fmt"
	"testing"
)

func PlusMinus(arr []float64) (float64, float64, float64) {

	positive, negative, nol := 0.0, 0.0, 0.0

	panjangIndex := len(arr)
	for i := 0; i < len(arr); i++ {
		if arr[i] > 0 {
			positive += 1
		} else if arr[i] < 0 {
			negative += 1
		} else {
			nol += 1
		}

	}
	sumPositive := positive / float64(panjangIndex)
	sumNegative := negative / float64(panjangIndex)
	sumBulat := nol / float64(panjangIndex)
	return sumPositive, sumNegative, sumBulat

}

func TestPlusMinus(t *testing.T) {

	var arr = []float64{-4, 3, -9, 0, 4, 1}
	positive, negative, bulat := PlusMinus(arr)

	fmt.Println(positive)
	fmt.Println(negative)
	fmt.Println(bulat)

	// fmt.Println(float64(positive) / float64(sum))
	// fmt.Println(float64(negative) / float64(sum))
	// fmt.Println(float64(nol) / float64(sum))

}
