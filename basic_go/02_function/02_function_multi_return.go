package main

import (
	"fmt"
)

func main() {
	/* -------------------------------------------------------------------------- */
	/*                   // * function multi return                               */
	/* -------------------------------------------------------------------------- */
	var side = 7
	var area, circumference = calculateSquare2(side)

	fmt.Println("luas persegi:", area)
	fmt.Println("keliling persegi:", circumference)
}

// * function with multiple return
func calculateSquare(s int) (int, int) {
	// luas
	var area = s * s

	// keliling
	var circumference = 4 * s

	return area, circumference
}

// * function with multiple predefined return
// kalo ini kita tidak usah mereturn variabelnya, karena sudah dideklarasikan di awal fungsi dan automatis akan direturn
func calculateSquare2(s int) (area, circumference int) {
	area = s * s
	circumference = 4 * s
	return
}
