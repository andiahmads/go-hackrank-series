package intreverse

import (
	"fmt"
	"math"
	"testing"
)

func Test(t *testing.T) {

	// new_int := 0
	number := 1534236469

	// result := ReverseInt(number)
	rev := reverse(number)
	// fmt.Println(result)
	fmt.Println(rev)

}

func ReverseInt(number int) int {

	//temporary int
	new_int := 0
	for number > 0 {
		reminder := number % 10

		new_int *= 10
		new_int += reminder

		number /= 10

		// fmt.Println(number)
	}
	return new_int
}

func reverse(x int) int {

	temp := 0

	maxINT := math.MaxInt

	test := 1534236469
	// fmt.Println(maxINT)
	fmt.Printf("ini int:%d \n", test)

	for x != 0 {

		c := x % 10
		temp = temp*10 + c

		x /= 10
		if x > maxINT {
			return 0
		}

	}
	return temp
}
