package modifyreverse_test

import (
	"fmt"
	"math"
	"testing"
)

func TestModifyReverse(t *testing.T) {
	world := "abcdefg"

	lenWorld := len(world)

	temporaryString := ""
	sisa := ""

	midWorld := 0

	if lenWorld%2 == 0 {
		midWorld = int(math.Floor(float64(lenWorld) / 2))
		temporaryString = world[:midWorld]
		sisa = world[midWorld:lenWorld]
		rev := ReverseModify(temporaryString)
		fmt.Printf("%s%s", rev, sisa)
	} else {
		midWorld = lenWorld / 2
		temporaryString = world[midWorld+1 : lenWorld]
		sisa = world[:midWorld+1]
		rev := ReverseModify(temporaryString)
		fmt.Printf("%s%s", rev, sisa)

	}
}

func ReverseModify(world string) string {

	temp := ""

	for i := 0; i < len(world); i++ {

		arr := string(world[i])

		temp = arr + temp

	}

	return temp

}
