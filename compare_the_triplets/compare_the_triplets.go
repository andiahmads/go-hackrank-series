package compare_the_triplets

func CompareTheTriplets(a, b [3]int) (int, int) {

	var cok, kocok int
	for i := 0; i < len(a); i++ {
		if a[i] > b[i] {
			cok += 1
		} else if a[i] < b[i] {
			kocok += 1
		}
	}

	return cok, kocok

}
