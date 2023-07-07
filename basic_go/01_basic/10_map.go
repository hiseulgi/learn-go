package main

import "fmt"

func main() {
	/* -------------------------------------------------------------------------- */
	/*                                 // * Map                                   */
	/* -------------------------------------------------------------------------- */
	// Map adalah tipe data key-value pair (dict kalo di python)
	// map mirip seperti slice, hanya indeksnya bisa kita tentukan sendiri

	// * penggunaan map
	var cow map[string]int
	cow = map[string]int{}

	cow["alex"] = 50
	cow["levi"] = 30

	fmt.Println(cow)
	fmt.Println(cow["alex"])
	fmt.Println()

	// * inisialisasi nilai map
	var fruits = map[string]int{
		"apple":   10,
		"grape":   5,
		"coconut": 8,
		"orange":  2,
		"cherry":  0,
	}

	// bisa juga pake gini
	// var fruits = map[string]int{}
	// var fruits = make(map[string]int)
	// var fruits = *new (map[string]int) // pake pointer

	// * iterasi item map for-range
	for key, val := range fruits {
		fmt.Println(key, ":", val)
	}
	fmt.Println()

	// * menghapus item map
	// dengan delete(map, "key")
	delete(fruits, "grape")

	for key, val := range fruits {
		fmt.Println(key, ":", val)
	}
	fmt.Println()

	// * deteksi keberadaan item pada key tertentu
	if value, isExist := fruits["salak"]; isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item is not exists")
	}
	fmt.Println()

	// * kombinasi slice & map
	// versi go lama
	// var students = []map[string]string{
	// 	map[string]string{"name": "asep", "gender": "m"},
	// 	map[string]string{"name": "udin", "gender": "m"},
	// 	map[string]string{"name": "septi", "gender": "f"},
	// }

	// for _, student := range students {
	// 	fmt.Println(student["name"], student["gender"])
	// }
	// fmt.Println()

	// versi go terbaru
	var students = []map[string]string{
		{"name": "asep", "gender": "m"},
		{"name": "udin", "gender": "m"},
		{"name": "septi", "gender": "f"},
	}
	fmt.Println(students)
	fmt.Println()
}
