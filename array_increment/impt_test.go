package arrayincrement

import (
	"fmt"
	"testing"
)

func IncrementArray(arr []int) {

	for i := range arr {
		arr[i]++
	}

}

func TestArrayIncrement(t *testing.T) {
	arr := []int{100, 200, 300}

	IncrementArray(arr)
	fmt.Println(arr)
}
