/* -------------------------------------------------------------------------- */
/*                            // * buffered channel                           */
/* -------------------------------------------------------------------------- */
// secara default channel bekerja dengan cara un-buffered pada memori
// ketika proses kirim data via channel dari sebuah goroutine, maka harus ada goroutine lain yang menerima dari channel tersebut
// sehingga bersifat blocking (saling tunggu)

// * buffered channel berbeda, kita dapat menentukan jumlah buffer-nya
// artinya berapa banyak data yang bisa dikirimkan secara bersamaan (multi-task)
// sehingga pengiriman akan berjalan async (tidak blocking)
// jadi secara ga langsung, dengan buffer kita bisa menentukan jumlah channel yang dapat bekerja dalam waktu yang sama
// ! namun perlu diingat, hanya pengiriman data saja yang bersifat async, untuk penerimaan datanya tetap sync

package main

import (
	"fmt"
)

func main() {
	// deklarasi var channel int dengan 3 buffer (4 data yang akan di-buffer)
	messages := make(chan int, 3) // buffer dimulai dari 0

	// make channel data receiver (async)
	// infinite loop menunggu pesan dari channel
	// hanya loop-nya saja yang berjalan secara async
	go func() {
		for {
			i := <-messages // untuk penerimaan data disini bersifat blocking (sync), jadi dia menunggu menerima datanya terlebih dahulu
			// goroutine diatas yang di-blocking, menunggu for dibawah memproses data baru
			fmt.Println("receive data", i)
		}
	}()

	// make channel data sender (sync, procedural)
	for i := 0; i < 6; i++ {
		fmt.Println("send data", i)
		messages <- i // blocking akan terjadi disini ketika channel buffer full
		// menunggu datanya diterima terlebih dahulu, baru berjalan kembali
	}
}

// jadi contoh diatas itu kita mengeksekusi terlebih dahulu channel data receiver terlebih dahulu pada goroutine
// kemudian baru mengeksekusi channel data sender secara prosedural

// * channel pada go bersifat fleksibel
// jadi kita bisa bisa secara bebas menentukan mana yang akan bersifat async dan bersifat sync
// tidak ada ketentuan yang kaku
