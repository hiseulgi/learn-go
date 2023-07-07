// * Time, Ticker, & Scheduler
// fungsi lain pada package time

package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// * time.Sleep()
	// digunakan untuk berhenti sejenak, bersifat blocking
	// fmt.Println("start")
	// time.Sleep(time.Second * 3)
	// fmt.Println("setelah 3 detik")

	// * scheduler
	// penjadwal pesan
	// ya kayak penjadwalan per berapa satuan waktu
	// for {
	// 	fmt.Println("hello anjay")
	// 	time.Sleep(1 * time.Second)
	// }

	// * fungsi time.NewTimer()
	// mengembalikan objek bertipe time.Timer yang memiliki properi c yang bertipe channel
	// cara kerja: setelah jeda waktu, sebuah data akan dikirimkan lewat channel c
	// dan nantinya harus diikuti dengan statement untuk penerimaan data dari channel c
	// timer := time.NewTimer(4 * time.Second) // inisialisasi objek Timer dan start timer 4 detik
	// fmt.Println("start")
	// <-timer.C // penerimaan data dari channel
	// fmt.Println("finish")

	// * fungsi time.AfterFunc()
	// menjalankan function setelah beberapa delay

	// ch := make(chan bool)
	// time.AfterFunc(time.Second*4, func() {
	// 	fmt.Println("expired")
	// 	ch <- true
	// }) // fungsi akan dieksekusi  setelah beberapa satuan waktu

	// fmt.Println("start")
	// <-ch // fungsi akan tereksekusi disini
	// fmt.Println("finish")

	// note: jika tidak ada serah data lewat channel, maka ekseksui AfterFunc() adalah async (tidak blocking)
	// jika ada serah terima data lewat channel, maka fungsi akan tetap berjalan async dan baru blocking ketika penerimaan data channel

	// * fungsi time.After()
	// seperti time.Sleep() tetapi fungsi ini akan mengembalikan data channel, sehingga perlu tanda <-
	// <-time.After(time.Second * 4)
	// fmt.Println("expired")

	/* -------------------------------------------------------------------------- */
	// * scheduler dengan Ticker
	// cara lain untuk penjadwalan, dimana dia bakal berjalan mengirim data setiap satuan waktu pada channel dan mengirim informasi date-time
	// done := make(chan bool)
	// ticker := time.NewTicker(time.Second) // inisialisasi objek ticker baru dengan delay 1 detik

	// // flag ticker done setelah 10 detik
	// go func() {
	// 	time.Sleep(time.Second * 10)
	// 	done <- true
	// }()

	// // menerima data terus menerus
	// for {
	// 	select {
	// 	// jika ada data channel done, maka ticker berhenti
	// 	case <-done:
	// 		ticker.Stop()
	// 		return
	// 	// menunggu pengiriman data ticker (1 detik)
	// 	case t := <-ticker.C:
	// 		fmt.Println("hello!", t)
	// 	}
	// }

	// * kombinasi timer dengan goroutine
	// contoh penerapan timer dan go, jadi kita bisa mengawasi program kita. apakah ada inputan dari user atau tidak. dan jika tidak ada, maka program akan memaksa exit dan selesai
	timeout := 5
	ch := make(chan bool)

	go timer(timeout, ch)
	go watcher(timeout, ch)

	input := ""
	fmt.Println("apakah saya ganteng? (y/n)")
	fmt.Scan(&input)

	if input == "y" {
		fmt.Println("tepat sekali wkwkwk")
	} else {
		fmt.Println("yah cacat wkwkwk")
	}
}

// digunakan sebagai timer dan dieksesusi oleh goroutine (async)
func timer(timeout int, ch chan<- bool) {
	time.AfterFunc(time.Duration(timeout)*time.Second, func() {
		ch <- true
	})
}

// digunakan sebagai pengamat timeout dan dieksekusi oleh goroutine (async)
func watcher(timeout int, ch <-chan bool) {
	<-ch
	fmt.Println("\ntimeout! no answer more than", timeout, "seconds")
	os.Exit(0)
}
