package main

import "fmt"

func main() {
	/* -------------------------------------------------------------------------- */
	/*                                 // * Slice                                 */
	/* -------------------------------------------------------------------------- */
	// -> reference elemen array, karena reference maka perubahan data akan berpengaruh pada slice lain yang alamat memorinya sama

	// * Inisialisasi Slice
	// mirip dengan array, tapi tanpa perlu jumlah elemen awal
	var fruits = []string{
		"apple",
		"grape",
		"orange",
		"melon",
	}
	fmt.Println(fruits)

	// * Operasi Slice
	// ini mirip2 ngambil data list di python
	var newFruitsA = fruits[0:2]
	var newFruitsB = fruits[2:]
	var newFruitsC = fruits[:2]

	fmt.Println()
	fmt.Println(newFruitsA)
	fmt.Println(newFruitsB)
	fmt.Println(newFruitsC)

	// * Mengubah data pada Slice
	// ini akan memengaruhi semuanya, karena satu referensi
	newFruitsA[0] = "starfruit"

	fmt.Println()
	fmt.Println(fruits)
	fmt.Println(newFruitsA)
	fmt.Println(newFruitsB)
	fmt.Println(newFruitsC)

	// * Fungsi len() dan cap()
	// len -> menghitung jumlah elemen slice yang ada
	// cap -> menghitung lebar dan kapasitas maksimum slice
	fmt.Println()
	fmt.Println(len(fruits)) // 4
	fmt.Println(cap(fruits)) // 4

	fmt.Println(len(newFruitsA)) // 2
	fmt.Println(cap(newFruitsA)) // 4, karena // [ada, ada, _, _]
	// slice setelah yang diambil dianggep kapasitas

	fmt.Println(len(newFruitsB)) // 2
	fmt.Println(cap(newFruitsB)) // 2, karena // _, _, [ada, ada]
	// slice yang sebelum yang diambil dianggep bukan kapasitas

	fmt.Println(len(newFruitsC)) // 2
	fmt.Println(cap(newFruitsC)) // 4

	// * Fungsi append()
	// menambah elemen pada slice, elemen tsb diposisikan pada indeks terakhir
	var iFruits = append(fruits, "mango") // membuat semua elemennya pada referensi baru, karena tidak ada cap yang kosong

	fmt.Println()
	fmt.Println(fruits)
	fmt.Println(iFruits)

	// pada append ini juga berlaku referensi
	// misal jika len == cap maka append akan membuat referensi baru
	// namun jika len < cap maka append akan membuatkannya ke referensi yang kosong
	var jFruits = fruits[:2]
	var kFruits = append(jFruits, "papaya") //  referensinya sama seperti jFruits dan fruits. karena masih ada cap yang kosong

	fmt.Println()
	fmt.Println(fruits)  // [starfruit grape papaya melon]
	fmt.Println(iFruits) // [starfruit orange papaya melon]
	fmt.Println(jFruits) // [starfruit grape]
	fmt.Println(kFruits) // [starfruit grape papaya]

	// * Fungsi copy()
	// copy(dst, src)
	// akan menyalin sesuai len pada var tujuan
	var dst = make([]int, 2)
	var src = []int{1, 2, 3, 4}
	var n = copy(dst, src)

	fmt.Println()
	fmt.Println(dst) // [1 2]
	fmt.Println(src) // [1 2 3 4]
	fmt.Println(n)   // 2

	// jika len var tujuan lebih besar dari len src, maka akan tetap menyalin namun pada index sesuai len src
	var dst2 = []int{1, 2, 3, 4, 5}
	var src2 = []int{6, 7}
	n = copy(dst2, src2)
	fmt.Println()
	fmt.Println(dst2) // [6 7 3 4 5]
	fmt.Println(src2) // [6 7]
	fmt.Println(n)    // 2

	// * Akses elemen slice dengan 3 indeks
	// slicing elemen sekaligus menentukan kapasitasnya
	fmt.Println()
	fmt.Println(fruits) // [starfruit grape papaya melon]

	var tempFruits = fruits[:2:2]
	fmt.Println(tempFruits)      // [starfruit grape]
	fmt.Println(len(tempFruits)) // 2
	fmt.Println(cap(tempFruits)) // 2
}
