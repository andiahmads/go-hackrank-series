package fizzbuzz

import (
	"fmt"
	"testing"
)

func TestFIzzBuzz(t *testing.T) {

	for i := 1; i < 15; i++ {

		if i%5 == 0 && i%3 == 0 {
			fmt.Println("fizzbuzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")

		} else if i%5 == 0 {
			fmt.Println("buzz")

		} else {
			fmt.Println(i)

		}

	}

}
