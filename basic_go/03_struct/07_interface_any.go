/* -------------------------------------------------------------------------- */
/*                               // * interface                               */
/* -------------------------------------------------------------------------- */
// interface -> kumpulan definisi method yang tidak memiliki isi
// merupakan tipe data, zero value-nya adalah nil

// * interface kosong (any)
// sebuah tipe data yang special. dapat diisi dengan semua jenis data mulai dari variabel biasa, array, map, slice, pointer, dan objek

package main

import (
	"fmt"
)

func main() {
	// mencoba tipe data interface
	var temp interface{}

	temp = "kids"
	fmt.Println(temp)

	temp = 10
	fmt.Println(temp)

	temp = []int{1, 2, 69}
	fmt.Println(temp)

	// selain interface{} kita bisa pake any
	var kosong any = 80
	fmt.Println(kosong)

	// interface kosong ini jika di-print akan menajadi string
	// namun ketika ingin menggunakan operasi aritmatika atau operasi pada variabel tertentu
	// kita perlu casting variable terlebih dahulu
	temp = 20
	// total := temp + kosong // error invalid operation
	total := temp.(int) + kosong.(int) // ga error
	fmt.Println(total)

	// * menyimpan pointer
	type person struct {
		name string
		age  int
	}

	var p1 any = &person{"andi", 31}
	fmt.Println(p1.(*person).name)

	// * gabungan slice, maps, dan interface
	// ini salah satu yang sering dipakai, contohnya pada encode/decode json
	// alternatif slice struct
	var students = []map[string]any{
		{"nama": "bagus", "umur": 20},
		{"nama": "andi", "umur": 23},
	}
	for _, student := range students {
		fmt.Println(student["nama"], "berumur", student["umur"])
	}

	// ini contoh tp rumit wkwkw
	var res = []any{
		map[string]any{
			"message": "ok",
			"status":  200,
			"data": []map[string]any{
				{"nama": "bagus"},
				{"umur": 18},
			},
		},
	}

	fmt.Println(res)
}
