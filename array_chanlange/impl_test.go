package arraychanlange

import (
	"fmt"
	"testing"
)

func TestArrayChanlge(t *testing.T) {

	arr := []int{4, 6, 23, 10, 1, 3}
	// arr := []int{4, 6, 23, 10, 1, 3}
	// arr := []int{3, 5, -1, 8, 12}
	// arr := []int{5, 7, 16, 1, 2}
	// arr := []int{3, 4, 1, 6}

	res := arrayChallenge(arr)

	fmt.Println(res)

}

/*
Mintalah fungsi ArrayChallenge(arr) mengambil array angka yang disimpan di arr dan mengembalikan string true
jika ada kombinasi angka dalam array (tidak termasuk angka terbesar) dapat ditambahkan hingga sama dengan angka terbesar dalam array,
jika tidak, kembalikan string palsu. Misalnya: jika arr berisi [4, 6, 23, 10, 1, 3] output harus mengembalikan true karena 4 + 6 + 10 + 3 = 23.
Array tidak akan kosong, tidak akan berisi semua elemen yang sama, dan mungkin berisi angka negatif.

*/

func arrayChallenge(arr []int) string {

	max := 0

	index := 0

	for i := 0; i < len(arr); i++ {

		if arr[i] >= max {
			max = arr[i]
			index = i
		}
	}

	removeaRR := append(arr[:index], arr[index+1:]...)

	fmt.Println(removeaRR)

	lenRemoveArr := len(removeaRR) - index + 1

	var count int

	for i := 0; i <= lenRemoveArr+1; i++ {
		count = 0

		batas := lenRemoveArr

		for j := i; j < batas; j++ {
			count += removeaRR[j] // 4, 6,10, 1, 3
		}

		if count == max {
			return "true"
		}
	}

	return "false"

}
