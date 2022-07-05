package stringpalindrom_test

import (
	"fmt"
	"testing"
)

func TestStringPalindrome(t *testing.T) {
	str := "abx"
	res := isPolindrome(str)
	fmt.Println(res)
}

func isPolindrome(str string) bool {
	for i := 0; i < len(str); i++ {
		// fmt.Print(string(str[i]))
		fmt.Print(i)

		j := len(str) - 1 - i

		/*
			3-1= 2 -0 = 2
			3-1= 2- 1 = 1
			3-1= 2- 2 = 0
		*/

		fmt.Println(j)
		if str[i] != str[j] {
			return false
		}
	}
	return true

}
