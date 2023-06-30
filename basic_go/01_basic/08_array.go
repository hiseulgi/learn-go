package main

import "fmt"

func main() {

	// * Array
	var names [3]string
	names[0] = "Asep"
	names[1] = "Bagus"
	names[2] = "Cendy"
	// names[3] = "Daril" // error karena melebihi batas jumlah array yang dibut

	fmt.Println(names[2])
	fmt.Println()

	// Inisialisasi nilai array di awal
	// var nilai = [3]int{80, 20, 6}

	// bisa juga kek gini (vertikal)
	var nilai = [3]int{
		69,
		40,
		70,
	}

	fmt.Println(len(nilai))
	fmt.Println(nilai)
	fmt.Println()

	// Inisialisasi array tanpa jumlah elemen
	// hanya untuk inisialisasi awal saja, tidak bisa fleksibel variabelnya (gabisa nambah value diluar jumlah awal)
	var grade = [...]int{
		1,
		1,
		2,
		3,
		7,
	}

	fmt.Println(len(grade))
	fmt.Println(grade)
	fmt.Println()

	// Array multidimensi
	var numbers1 = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	fmt.Println(len(numbers1))
	fmt.Println(numbers1)
	fmt.Println(numbers1[0][2])
	fmt.Println()

	// Penggunaan For pada Array
	// for biasa
	var fruits = [...]string{
		"apel",
		"jeruk",
		"ceri",
	}

	fmt.Println("for biasa")
	for i := 0; i < len(fruits); i++ {
		fmt.Printf("%s ", fruits[i])
	}
	fmt.Println()

	// for - range
	fmt.Println("for range")
	for _, fruit := range fruits { // kita gapake indeksnya jadi variabel awal pake _
		fmt.Printf("%s ", fruit)
	}
	fmt.Println()
	fmt.Println()

	// Inisialisasi Array dengan Make
	var classes = make([]int, 2)
	classes[0] = 25
	classes[1] = 50
	fmt.Println(classes)

}
