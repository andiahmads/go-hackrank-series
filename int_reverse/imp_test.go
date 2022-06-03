package intreverse

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {

	// new_int := 0
	number := 153
	// number := 153

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
	prev := 0

	for x != 0 {

		c := x % 10

		x /= 10

		temp = temp*10 + c

		if temp-c/10 != prev {
			return 0
		}
		prev = temp

	}
	return temp
}
