package anagram

import (
	"fmt"
	"testing"
)

func TestAnagram(t *testing.T) {
	res := AreAnagram("abc", "abc")
	fmt.Println(res)
}

func AreAnagram(s1, s2 string) bool {

	if len(s1) != len(s2) {
		fmt.Println("string not anagram")
	}

	hash := make(map[string]int)

	for _, r := range s1 {
		j := hash[string(r)]

		if j == 0 {
			hash[string(r)] = 1
		} else {
			hash[string(r)] = j + 1
		}
	}

	for _, r := range s2 {
		j := hash[string(r)]

		if j == 0 {
			hash[string(r)] = 1
		} else {
			hash[string(r)] = j + 1
		}
	}

	var isAnagram bool = true

	for _, value := range hash {
		if value%2 != 0 {
			isAnagram = false
			break
		}
	}
	return isAnagram


}
