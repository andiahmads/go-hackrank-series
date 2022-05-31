package polindromenumber_test

import (
	"fmt"
	"testing"
)

func TestPolindromenumber(t *testing.T) {
	res := isPolindrome(121)
	fmt.Println(res)
}

func isPolindrome(x int) bool {

	original := x

	reminder := 0

	reversed := 0

	if original < 0 {
		original *= -1
	}

	for x != 0 {

		reminder = x % 10
		// fmt.Printf("reminder:%v \n", reminder)
		reversed = reversed*10 + reminder
		// fmt.Printf("no:%v \n", reversed)
		x /= 10

	}

	fmt.Printf("original:%v \n", original)
	fmt.Printf("reversed:%v \n", reversed)

	if original == reversed {
		return true
	} else {
		return false
	}

	// sum = (sum * 10) + r

	// n = n / 10

	// if temp == sum {
	// 	return true
	// } else {
	// 	return false
	// }

}
