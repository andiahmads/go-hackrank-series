package teknusajaya_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

// func mindDiff(arr []int) int {

// 	sum := 0.0

// 	sort := sort.IntSlice(arr)

// 	fmt.Println(sort)

// 	for i := 0; i < len(sort)-1; i++ {
// 		fmt.Println(arr[i], arr[i+1], arr[i]-arr[i+1])
// 		sum += math.Abs(float64(sort[i]) - float64(sort[i+1]))
// 	}

// 	return int(sum)

// 	n := int64(123)

// 	fmt.Println(strconv.FormatInt(n, 2)) // 1111011

// }

func TestMinDiff(t *testing.T) {

	// arr := []int{5, 1, 3, 7, 3}

	res := changeAds(50)
	fmt.Println(res)

	// res := mindDiff(arr)

	// fmt.Println(res)
}

func changeAds(base10 int32) int32 {

	s := strconv.FormatInt(int64(base10), 2)

	tes := strings.Split(s, "")

	b := []int{1, 2, 4, 8, 16, 32, 64}
	res := 0
	length := len(tes)

	for _, val := range tes {
		if val == "0" {
			res += b[length-1]
		}
		length -= 1
		fmt.Println(length)

	}

	return int32(res)

}
