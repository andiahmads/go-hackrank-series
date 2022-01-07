package compare_the_triplets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompareTheTriplets(t *testing.T) {

	var andi = [3]int{3, 5, 3}
	var ahmad = [3]int{1, 4, 5}

	a, b := CompareTheTriplets(andi, ahmad)
	result := "2 1"

	assert.Equal(t, a, b, result)

}
