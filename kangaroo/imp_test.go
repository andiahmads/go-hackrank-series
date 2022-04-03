package kangaroo

import (
	"fmt"
	"testing"
)

/*
x1 = posisi awal kangaroo 1 = 0
x2 = posisi awal kangaroo 2 = 2

x1 < x2

v1 = kecepatan melompat kangaroo 1 = 5
v2 = kecepatan melompat kangaroo 2 = 3


0,1,2,3,4,5,6,7,8,9,10

*/

func TestKangaroo(t *testing.T) {
	res := Kangaroo(0, 3, 4, 2)
	fmt.Println(res)
}

func Kangaroo(x1, x2, v1, v2 int) string {

	result := ""

	for x1 <= x2 {
		x1 += v1
		x2 += v2
		if x1 == x2 && v1 == v2 {
			result = "YES"
		} else if x1 == x2 && v1 != v2 {
			result = "NO"
		}
	}
	return result

}
