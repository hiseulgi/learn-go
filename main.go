package main

import (
	"fmt"
)

func main() {
	/* --------------------------- // ! BATAS COMMENT --------------------------- */
	// /* -------------------------------------------------------------------------- */
	// /*                             // * 1. Hello World                            */
	// /* -------------------------------------------------------------------------- */
	// // Print Hello World! to the console
	// /*
	// 	Ini Multiline Komentar
	// */
	// fmt.Println("Hello World!")

	// /* -------------------------------------------------------------------------- */
	// /*                              // * 2. Variable                              */
	// /* -------------------------------------------------------------------------- */
	// // Variable dengan Manifest Typing / Tipe Data
	// var firstName string = "Bagus"
	// var lastName string = "Adi"

	// fmt.Printf("Halo %s %s!\n", firstName, lastName)

	// // Variable tanpa Tipe Data
	// // pakenya := (hanya bisa dipake di blok fungsu)
	// // var fullName = "Asep Sugeng" // Ini sama aja kek di bawha
	// fullName := "Asep Sugeng"

	// fmt.Println(fullName)

	// // Multi Variable
	// var one, two, three string
	// one, two, three = "satu", "dua", "tiga"

	// cat, dog := 1, false

	// fmt.Println(one, two, three, cat, dog)

	// // Reserved Variable _
	// // Dalam Go, semua var harus dipakai. Maka ada _ untuk menampung nilai yg tidak dipakai
	// // Semua value yang masuk var _ akan hilang, biasanya dipakai untuk nilai return pada fungsi yang tidak digunakan
	// // var _ tidak dapat digunakan
	// username, _ := "nanang", "kusnandar"
	// fmt.Println(username)
	// // fmt.Println(_)

	// Variable dengan new
	// new akan membuat pointer ke tipe data tertentu
	// isi var dengan new adalah alamat dari variabel
	// untuk melihat nilainya menggunakan deference (*)
	// name := new(string)
	// fmt.Println(name)  //0xc000014270
	// fmt.Println(*name) // ""
	/* --------------------------- // ! BATAS COMMENT --------------------------- */

	/* -------------------------------------------------------------------------- */
	/*                             // * 3. TIPE DATA                              */
	/* -------------------------------------------------------------------------- */
	// Numerik Non-Desimal
	/*
		tipe data / ekuivalennya
		uint8 / byte
		uint16 / uint
		uint32 / uint
		uint64
		int8
		int16
		int32 / int / rune
		int64 / int
	*/
	var positiveNumber uint8 = 89
	var negativeNumber = -12692833
	// secara otomatis compiler akan menentukan tipe var yang cocok pada var negativeNumber jika deklarasi tanpa tipe data

	fmt.Printf("bilangan positif= %d\n", positiveNumber)
	fmt.Printf("bilangan negatif= %d\n", negativeNumber)

	// Numerik Desimal
	/*
		float32
		float64
	*/

	var decimalNumber float32 = 11.75
	fmt.Printf("bilangan desimal= %f\n", decimalNumber)
	fmt.Printf("bilangan desimal= %.2f\n", decimalNumber)

	// Boolean
	var isOn = true
	fmt.Printf("isOn? %t \n", isOn)

	// String
	// bisa juga dengan backticks ``
	var message = `Nama saya "Bagus".
Saya senang belajar Go.`

	fmt.Println(message)

	// Nil (semacam null value)
	// Nil berbeda dengan zero value, dimana zero value adalah nilai default dari variable yang tanpa nilai
	// sedangkan Nil adalah nilai kosong, hanya dapat digunakan pada pointer, tipe data fungsi, slice, map, channel, interface kosong/any
}
