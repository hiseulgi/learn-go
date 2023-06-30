/* -------------------------------------------------------------------------- */
/*                              // * defer & exit                             */
/* -------------------------------------------------------------------------- */
// defer -> mengakhirkan eksekusi statement tepat sebelum blok selesai
// exit menghentikan program secara paksa (exit terminal)

package main

import (
	"fmt"
	"os"
)

// * defer pada func
func orderFoodPlease(menuName string) {
	defer fmt.Println("thanks sudah pesan!")

	fmt.Println("oke kamu akan pesan ya!")

	if menuName == "burger" {
		fmt.Println("keren makan burger")
	}

}

func main() {
	// defer fmt.Println("Aku yang pertama harusnya") // dieksesusi terakhir sebelum blok main selesai
	// fmt.Println("Hai ini adalaha bagian kedua")

	// * defer pada func
	// orderFoodPlease("pizza")
	// orderFoodPlease("burger")

	// * defer pada closuer IIFE
	fmt.Println("saya yang 1")

	// ini bakal dieksekusi terlebih dahulu karena ia ada di dalam scope clousre IIFE
	func() {
		defer fmt.Println("saya yang 3")
	}()
	fmt.Println("saya yang 2")

	// * os.Exit() untuk menghentikan program secara paksa
	// pada bagian ini, defer tidak akan dieksekusi karena program sudah keluar terlebih dahulu pada baris os.Exit(1)
	defer fmt.Println("saya harusnya terakhir")
	fmt.Println("sebelum exit bosku")
	os.Exit(1)
	fmt.Println("anjay")
}
