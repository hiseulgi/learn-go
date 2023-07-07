/* -------------------------------------------------------------------------- */
/*                               // * goroutine                               */
/* -------------------------------------------------------------------------- */
// - mirip seperti thread namun lebih kecil (mini thread)
// - berjalan secara async, jadi tidak saling tunggu
// - bagian penting dalam concurrent programming di Go
// - eksekusinya dijalankan di multi core processor, jadi kita dapat menentukan berapa banyak core yang aktif

// concurrency beda dengan paralel
// concurrency adalah komposisi dari sebuah proses dan merupakan struktur
// paralel adalah eksekusi banyak proses secara bersamaan dan bagaimana eksekusinya berlangsung

// jadi konkurensi adalah bagaimana kita mengatur sebuah komposisi atau struktur program agar dapat berjalan dalam wsatu waktu yang sama

package main

import (
	"fmt"
	"runtime"
	"time"
)

func print(max int, message string) {
	for i := 0; i < max; i++ {
		fmt.Println(i, ":", message)
	}
}

func main() {
	fmt.Println(runtime.NumCPU()) // ini cuma buat liat max cpu pada laptop

	// set max processor pada runtime goroutine
	// runtime.GOMAXPROCS(2)

	// untuk menandakan goroutine, tambahakan sintaks go di awal
	go print(5, "anjay")
	print(5, "apa kabss")

	go print(5, "sokin")

	// blocking (agar dapat melihat jalannya goroutine)
	// var input string
	// fmt.Scanln(&input)

	/* -------------------------------------------------------------------------- */
	// ini juga bisa buat test, hasilnya gak berurutan
	go func() { fmt.Println("hi I'm a goroutine1!") }()
	go func() { fmt.Println("hi I'm a goroutine2!") }()
	go func() { fmt.Println("hi I'm a goroutine3!") }()
	fmt.Println("Hello, we are testing goroutines in Go!")
	time.Sleep(time.Second)
}

// output program diatas adalah
// go print(5, "anjay") akan tereksekusi lebih lama dibanding print pada scope main yang mana tidak saling tunggu
// hasil output dari program juga bakal berbeda-beda karena kita menggunakan 2 procesor
