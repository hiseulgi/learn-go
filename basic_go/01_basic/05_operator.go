package main

import "fmt"

func main() {
	/* -------------------------------------------------------------------------- */
	/*                              // * 5. Operator                              */
	/* -------------------------------------------------------------------------- */
	// * A. Operator Aritmatika
	// sama seperti bahasa lain
	// +-*/%
	var value = (((2+6)%3)*4 - 2) / 3
	fmt.Println(value)

	// * B. Operator Perbandingan
	// ==, !=, <, <=, >, >=
	var isEqual = (value != 3)
	fmt.Printf("nilai %d (%t) \n", value, isEqual)

	// * C. Operator Logika
	// &&, ||, !
	var left = false
	var right = true

	fmt.Println(left && right)
	fmt.Println(left || right)
	fmt.Println(!left)
}
