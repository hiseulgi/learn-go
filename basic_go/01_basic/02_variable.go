package main

import "fmt"

func main() {
	/* -------------------------------------------------------------------------- */
	/*                              // * 2. Variable                              */
	/* -------------------------------------------------------------------------- */
	// Variable dengan Manifest Typing / Tipe Data
	var firstName string = "Bagus"
	var lastName string = "Adi"

	fmt.Printf("Halo %s %s!\n", firstName, lastName)

	// Variable tanpa Tipe Data
	// pakenya := (hanya bisa dipake di blok fungsu)
	// var fullName = "Asep Sugeng" // Ini sama aja kek di bawha
	fullName := "Asep Sugeng"

	fmt.Println(fullName)

	// Multi Variable
	var one, two, three string
	one, two, three = "satu", "dua", "tiga"

	cat, dog := 1, false

	fmt.Println(one, two, three, cat, dog)

	// Reserved Variable _
	// Dalam Go, semua var harus dipakai. Maka ada _ untuk menampung nilai yg tidak dipakai
	// Semua value yang masuk var _ akan hilang, biasanya dipakai untuk nilai return pada fungsi yang tidak digunakan
	// var _ tidak dapat digunakan
	username, _ := "nanang", "kusnandar"
	fmt.Println(username)
	// fmt.Println(_)

	// Variable dengan new
	// new akan membuat pointer ke tipe data tertentu
	// isi var dengan new adalah alamat dari variabel
	// untuk melihat nilainya menggunakan deference (*)
	name := new(string)
	fmt.Println(name)  //0xc000014270
	fmt.Println(*name) // ""

}
