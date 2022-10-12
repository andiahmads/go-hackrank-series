package coin

import "testing"

func TestCoint(t *testing.T) {

}

func solve(c []int32) int32 {

	var result int32 = 1

	var posibleCard int32

	var count int32 = 0

	for i := 0; i < len(c); i++ {
		if c[i] <= int32(i) {
			count++
		}
	}

	for i := 0; i < len(c); i++ {
		posibleCard = count

		if posibleCard == 0 {
			return 0
		} else {
			result *= posibleCard - int32(i)
		}

	}

	return result

}
