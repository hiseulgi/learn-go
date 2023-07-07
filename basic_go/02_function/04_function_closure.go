package main

import "fmt"

func main() {
	/* -------------------------------------------------------------------------- */
	/*                             // * fungsi closure                            */
	/* -------------------------------------------------------------------------- */
	// closure -> fungsi yang bisa disimpan dalam variabel
	// bisa dikatakan fungsi tanpa nama (anonymous function)
	// biasa digunakan untuk membungkus proses yang sekali pakai

	// * closure disimpan sebagai variabel
	// ex: mengambil nilai terendah dan tertinggi pada array
	getMinMax := func(n []int) (int, int) {
		var min, max int
		for i, e := range n {
			switch {
			case i == 0:
				min, max = e, e
			case max < e:
				max = e
			case min > e:
				min = e
			}
		}
		return min, max
	}

	numbers := []int{2, 3, 1, 0, 9, 10, 5, 7, 11}
	min, max := getMinMax(numbers) // memanggil fungsi
	fmt.Printf("data : %v\nmin : %v\nmax: %v\n", numbers, min, max)
	fmt.Println()

	// * immediately-invoked function expression (iife)
	// langsung dieksekusi saat deklarasinya
	// ex: filter nilai pada array
	numbers = append(numbers, 4, 5, 6, 9, 9, 0, 0)
	newNumbers := func(min int) []int {
		var r []int
		for _, number := range numbers {
			if number >= min {
				r = append(r, number)
			}
		}
		return r
	}(3) // fungsi langsung dieksekusi karena dibungkus ini (3)
	fmt.Printf("input: %v\noutput: %v\n", numbers, newNumbers)
	fmt.Println()

	// * closure sebagai return value
	// jadi dengan ini kita bisa membuat return value sebuah fungsi
	// fungsi tersebut dapat dijalankan lagi untuk mendapatkan output yang dikemablikannya
	len, getNumbers := findMax(numbers, 8) // ini mengembalikan int dan fungsi
	maxNumbers := getNumbers()             // run fungsi untuk mendapatkan return value
	fmt.Printf("input: %v\noutput: %v\nlen: %v\n", numbers, maxNumbers, len)
}

// * closure sebagai return value
func findMax(numbers []int, max int) (int, func() []int) {
	var r []int
	for _, number := range numbers {
		if number <= max {
			r = append(r, number)
		}
	}
	return len(r), func() []int {
		return r
	}
}
