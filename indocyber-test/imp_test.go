package indocybertest

import (
	"fmt"
	"testing"
)

//buat 1 function rangking

// bentuk array standar rangking

// dari batas atas sampai batas bawah

// contoh [100,80,50,20,10]

// nilai yg didapatkan oleh user itu berbentuk array
// user mendapatkan 3 nilai
// [55,80,105]

// output [2,1,1]

func Rangkin(arr []int, n1, n2, n3 int) []int {

	// min := 0
	// fmt.Println(max)
	nilai := []int{}
	for i := 0; i < len(arr); i++ {
		if n1-5 == arr[i] || n2 == arr[i] || n3-5 == arr[i] {
			if i == 0 {
				i = 1
			}
			nilai = append(nilai, i)
		}

	}

	return nilai

}

func TestRangking(t *testing.T) {

	arr := []int{100, 80, 50, 20, 10}

	res := Rangkin(arr, 55, 80, 105)

	fmt.Println(res)

}
