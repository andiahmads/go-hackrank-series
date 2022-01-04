package main

import (
	"fmt"
	"github/andiahmads/go-hackrank-series/simple_array_sum.go"
)

func main() {
	var data = [6]int{1, 2, 3, 4, 10, 11}
	var res int
	res = simple_array_sum.ArraySum(data, 6)
	fmt.Printf("Final result is: %d ", res)
}
