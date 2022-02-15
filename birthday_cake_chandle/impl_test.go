package birthdaycakechandle

import (
	"fmt"
	"testing"
)

func TestBirthday(t *testing.T) {
	arr := []int{3, 2, 1, 3}

	total := Lilin(arr)

	fmt.Println(total)

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
