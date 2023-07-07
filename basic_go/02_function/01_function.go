package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	/* -------------------------------------------------------------------------- */
	/*                                // * function                               */
	/* -------------------------------------------------------------------------- */
	// ayo mulai kurangi dry (dont repeat yourself)
	// * fungsi tanpa return value
	var names = []string{"nanang", "kusnandar"}
	printMessage("halo", names)

	// * fungsi dengan return value
	var randomValue int = randomWithRange(1, 10)
	fmt.Println("random value:", randomValue)

	// * return bisa dipakai untuk keluar dari function
	testLoop(randomValue)
}

// * fungsi tanpa return value
func printMessage(message string, arr []string) {
	var nameString = strings.Join(arr, " ")
	fmt.Println(message, nameString)
}

// * fungsi dengan return value
// var randomizer = rand.New(rand.NewSource(6969)) // yang di dalem itu seed-nya
var randomizer = rand.New(rand.NewSource(time.Now().Unix()))

// tipe data return value ditulis di belakang sebelum tanda {
// pada parameter dibawah juga karena bertipe data sama, dapat dituliskan seperti itu (nama_var1, nama_var2 tipe_data)
// func nameOfFunc(paramA, paramB type, paramC type) returnType
func randomWithRange(min, max int) int {
	return randomizer.Int()%(max-min+1) + min
}

// * return bisa dipakai untuk keluar dari function
func testLoop(input int) {
	for i := 0; i < 10; i++ {
		if i == input {
			fmt.Println("kamu cacat")
			return
		}
		fmt.Println(i)
	}
}
