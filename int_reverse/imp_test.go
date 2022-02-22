package intreverse

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {

	// new_int := 0
	number := 41234

	result := ReverseInt(number)
	fmt.Println(result)

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
