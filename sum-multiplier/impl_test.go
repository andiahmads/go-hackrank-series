package summultiplier_test

import (
	"fmt"
	"testing"
)

/*

Mintalah fungsi SumMultiplier(arr) mengambil larik angka yang disimpan dalam arr
dan mengembalikan string true jika ada dua angka yang dapat dikalikan
sehingga jawabannya lebih besar dari dua kali lipat jumlah semua elemen dalam larik.
Jika tidak, kembalikan string false. Contoh: jika arr adalah [2, 5, 6, -6, 16, 2, 3, 6, 5, 3]
maka jumlah semua elemen ini adalah 42 dan menggandakannya adalah 84. Ada dua elemen dalam array ,
16 * 6 = 96 dan 96 lebih besar dari 84, jadi program Anda harus mengembalikan string yang benar.

*/

func TestSumMultiplier(t *testing.T) {

	arr := []int{2, 2, 2, 2, 4, 1}

	res := SumMultiplier(arr)

	fmt.Println(res)

}

func SumMultiplier(arr []int) string {

	temp := 0

	for i := 0; i < len(arr); i++ {
		temp += arr[i] * 2

	}

	for x := 0; x < len(arr)-1; x++ {

		for y := x + 1; y < len(arr); y++ {

			if arr[x]*arr[y] > temp {
				return "true"
			}

		}

	}

	return "false"

}
