package circulararrayrotation

import (
	"fmt"
	"testing"
)

func TestCircularArrayRotate(t *testing.T) {
	// 1,2,3
	// 3,1,2
	// 2,3,1
	// 1,2,3

	var fruits = []int{1, 2, 3}

	rotate(&fruits, 3)

	fmt.Println(fruits)

}

func rotate(arr *[]int, totalRotate int) {
	x, b := (*arr)[:(len(*arr)-totalRotate)], (*arr)[(len(*arr)-totalRotate):]
	fmt.Println(b)
	*arr = append(b, x...)
}

func leftRotation(a []int, size, rotation int) []int {

	var newArrays []int

	for i := 0; i < rotation; i++ {
		newArrays = a[1:size]
		newArrays = append(newArrays, a[0])
		a = newArrays
	}

	return a

}

func TestRotation(t *testing.T) {
	a := []int{1, 2, 3}
	fmt.Println("before rotation = ", a)
	rotation := 1

	fmt.Printf("output: %+v\n", leftRotation(a, len(a), rotation))
}
