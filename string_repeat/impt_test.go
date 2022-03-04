package stringrepeat

import (
	"fmt"
	"testing"
)

func RepeatString(s string, n int64) int64 {
	inputLength := int64(len(s))

	fmt.Println("inputLength = ", inputLength)

	numberOfSubstring := int64(n / inputLength)
	fmt.Println("numberOfSubstring = ", numberOfSubstring)

	totalCharA := CountCharA(s)

	numberOfAs := numberOfSubstring * totalCharA
	fmt.Println("numberOfAs = ", numberOfAs)

	reminder := n % inputLength
	fmt.Println("reminder = ", reminder)

	for i := int64(0); i < reminder; i++ {
		fmt.Println(string(s[i]))
		if s[i] == 'a' {
			numberOfAs++
		}

	}
	return numberOfAs
}

// abaaabaaaba

func CountCharA(s string) int64 {
	var count int64
	for i := 0; i < len(s); i++ {
		if s[i] == 'a' {
			count++
		}
	}
	return count
}

func TestStringRepeat(t *testing.T) {
	res := RepeatString("aba", 1)

	fmt.Println(res)
}
