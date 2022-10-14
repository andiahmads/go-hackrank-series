package relativerank

import (
	"fmt"
	"sort"
	"strconv"
	"testing"
)

func findRelativeRanks(score []int) []string {

	mapping := make(map[int]int)

	medal := make([]string, len(score))

	for i := range score {
		mapping[score[i]] = i
	}

	sort.Ints(score)

	temp := 4

	for i := len(score) - 1; i >= 0; i-- {
		switch i {
		case len(score) - 1:
			medal[mapping[score[i]]] = "Gold Medal"
			break

		case len(score) - 2:
			medal[mapping[score[i]]] = "Silver Medal"
			break

		case len(score) - 3:
			medal[mapping[score[i]]] = "Bronze Medal"
			break
		default:
			medal[mapping[score[i]]] = strconv.Itoa(temp)
			temp++
		}
	}

	return medal

}

func TestRelativeRank(t *testing.T) {
	arr := []int{5, 4, 3, 2, 1}

	res := findRelativeRanks(arr)

	fmt.Println(res)
}
