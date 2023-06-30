package main

import (
	"fmt"
)

func main() {
	/* -------------------------------------------------------------------------- */
	/*                            // * fungsi variadic                            */
	/* -------------------------------------------------------------------------- */
	// variadic function adalah fungsi yang memiliki parameter dengan tipe data dan nilainya tak terbatas (bisa berapa saja)
	// deklarasinya pake (varName ...type)
	numbers := []int{9, 1, 2, 5, 4, 7, 2, 3}
	average := calculateAverage(2, 3, 4, 2, 4, 3, 1, 5, 6, 8)
	average2 := calculateAverage(numbers...) // untuk variadic func ini bisa dimasukan slice juga

	msg1 := fmt.Sprintf("rata-rata: %.2f", average) // Sprintf -> digunakan untuk memformat output dan dimasukan ke variabel tipe string
	msg2 := fmt.Sprintf("rata-rata: %.2f", average2)
	fmt.Println(msg1)
	fmt.Println(msg2)
}

func calculateAverage(numbers ...int) float64 {
	total := 0
	for _, number := range numbers {
		total += number
	}

	average := float64(total) / float64(len(numbers))
	return average
}
