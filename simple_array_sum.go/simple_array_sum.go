package simple_array_sum

func ArraySum(arr [6]int, size int) int {
	var jumlah_hasil int

	for i := 0; i < size; i++ {
		jumlah_hasil += arr[i]
	}
	return jumlah_hasil

}
