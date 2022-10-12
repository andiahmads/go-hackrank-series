package birthdaycakechandle

import (
	"fmt"
	"testing"
)

func TestBirthday(t *testing.T) {
	arr := []int{3, 2, 1, 3, 2, 1, 1}

	// total := Lilin(arr)

	res := ChandleTwo(arr)

	// fmt.Println(total)
	fmt.Println(res)

}

func Lilin(chandle []int) int {

	var visitedElement = make(map[int]int)

	total := 0
	for i := 0; i < len(chandle); i++ {

		_, isExist := visitedElement[chandle[i]]
		if isExist {
			total = visitedElement[chandle[i]] + 1
		} else {
			visitedElement[chandle[i]] = 1

		}
	}
	return total
}

func ChandleTwo(chandle []int) int {

	// var visited = make(map[int]int)

	max := 0

	for i := 0; i < len(chandle); i++ {

		if chandle[i] >= max {
			max = chandle[i]
		}

	}

	res := 0
	for i := 0; i < len(chandle); i++ {
		if chandle[i] >= max {
			res++
		}
	}

	return 0

}
