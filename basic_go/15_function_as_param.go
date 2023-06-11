package main

import (
	"fmt"
	"strings"
)

// * 3. alias skema closure
// karena tipe parameter closure terlalu panjang, kita dapat memakai alias
// alias ini secara tidak langsung kita membuat semacam tipe data baru untuk sebuah fungsi
type FilterCallback func(string) bool

// * 1. fungsi filter, dengan parameter callback func
// callback func digunakan sbg filter yang digunakan dalam bentuk fungsi
// callback ini bertipe func dengan return value bool
// func filter(data []string, callback func(string) bool) []string {
func filter(data []string, callback FilterCallback) []string {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}

	return result
}

// * 2. main func
func main() {
	/* -------------------------------------------------------------------------- */
	/*                        // * fungsi sebagai parameter                       */
	/* -------------------------------------------------------------------------- */
	// init slice
	names := []string{"sugab", "adi", "riyan", "zia", "bagus", "david"}

	// mencari nama dengan 5 huruf
	namesLen5 := filter(names, func(s string) bool {
		return len(s) == 5
	})

	// mencari nama yang ada huruf i
	namesContainI := filter(names, func(s string) bool {
		return strings.Contains(s, "i")
	})

	fmt.Println("data asli:", names)
	fmt.Println("data dengan 5 huruf:", namesLen5)
	fmt.Println("data yang ada huruf i:", namesContainI)
}
