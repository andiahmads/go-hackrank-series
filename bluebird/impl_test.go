package bluebird

import (
	"fmt"
	"testing"
)

func TestBlueBird(t *testing.T) {
	arr := []int{8, 5, 9, 10, 2, 90}
	res := Bluebird(arr)
	fmt.Println(res)

}

func Bluebird(arr []int) int {

	terang := 0
	gedung_tarang := 0
	for _, v := range arr {

		if v > terang {
			terang = v
			gedung_tarang++
		}

	}
	return gedung_tarang
}
