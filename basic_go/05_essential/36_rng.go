// * Random Number Generator (RNG)
// penghasil data angka acak
// seed -> titk mulai (starting point) pada RNG
// di Golang ada package math/rand yang mengadopsi PRNG/pseudo-random number generator

package main

import (
	"fmt"
	"math/rand"
	"time"
)

var randomizer = rand.New(rand.NewSource(10))
var randomizerUnique = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	// seed: 10
	fmt.Println("angka ke-1:", randomizer.Int())
	fmt.Println("angka ke-2:", randomizer.Int())
	fmt.Println("angka ke-3:", randomizer.Int())

	// unique seed
	// dengan menggunakan timestamp unix nano
	fmt.Println("angka ke-1:", randomizerUnique.Int())
	fmt.Println("angka ke-2:", randomizerUnique.Int())
	fmt.Println("angka ke-3:", randomizerUnique.Int())

	// tipe data lain
	fmt.Println(randomizer.Float64()) //float
	fmt.Println(randomizer.Uint64())  //unsigned int
	fmt.Println(randomizer.Intn(10))  //int dengan batas

	// random data string
	fmt.Println(randomString(10))
	fmt.Println(randomString(10))
	fmt.Println(randomString(10))
}

// randomizer string
// algoritma dia random int number dengan batas total char inputan, kemudian random pada slice of char tersebut
func randomString(length int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[randomizer.Intn(len(letters))]
	}

	return string(b)
}
