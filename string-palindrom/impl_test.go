package stringpalindrom_test

import (
	"fmt"
	"strings"
	"testing"
)

func TestStringPalindrome(t *testing.T) {
	str := "Was it a car or a cat I saw?"

	res := isPalindrome2(str)

	fmt.Println(res)

}

func isPalindrome2(str string) bool {

	str = strings.ToLower(str)
	str = strings.ReplaceAll(str, " ", "")
	str = strings.ReplaceAll(str, "?", "")

	fmt.Println(str)
	for i := 0; i < len(str); i++ {

		j := len(str) - i - 1

		if str[i] == str[j] {
			return true
		}

	}

	return false

}

func isPalindrome(a string) bool {

	n := len(a)

	var st map[string]interface{}

	for i := 0; i < n; i++ {
		st[""] = string(a[i])
	}

	var check bool

	for _, val := range st {
		low := 0
		high := n - 1

		var flag bool

		for i := 0; i < n; i++ {
			if a[low] == a[high] {

				low++
				high--

			} else {
				if a[low] == val {
					low++
				} else if a[high] == val {

					high--
				} else {
					flag = false
					break
				}
			}
		}

		if flag == true {
			check = true
			break
		}
	}

	if check {
		return true
	} else {
		return false
	}

}
