package main

import "fmt"

func main() {
	/* -------------------------------------------------------------------------- */
	/*                              // * 4. Konstanta                             */
	/* -------------------------------------------------------------------------- */
	// * Konstanta
	const firstName string = "Bagus"
	const lastName = "Adi"
	// firstName = "Sugab" // Bakal Error karena const tidak bisa diubah

	fmt.Println("Namaku adalah", firstName, lastName)

	// * Multi Konstanta
	const (
		pi            = 3.14
		square        = "kotak"
		isToday bool  = true
		num     uint8 = 14
		angka         // dup value from above (num)
	)

	// bisa juga gini nih hehe
	const satu, dua = 1, 2
	fmt.Println(num, angka)

}
