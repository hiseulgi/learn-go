// * WaitGroup
// sync.WaitGroup -> sinkronisasi goroutine
// berbeda dengan channel, WaitGroup digunakan untuk maintain goroutine, dan lebih efektif dibanding channel

// ! warning
// sebenarnya fungsi keduanya berbeda
// channel untuk keperluan sharing antar goroutine
// WaitGroup untuk sinkronisasi goroutine

package main

import (
	"fmt"
	"sync"
)

// func to print using WaitGroup
func doPrint(wg *sync.WaitGroup, message string) {
	// 4. menandakan goroutine sudah selesai
	defer wg.Done()
	fmt.Println(message)
}

func main() {
	// 1. make new waitgroup variable
	wg := sync.WaitGroup{}

	// loop all iterate
	for i := 0; i < 5; i++ {
		data := fmt.Sprintf("data %d", i)

		// 2. menambah goroutine yang sedang diproses sebanyak 1
		// jadi totalnya akan ada 5 goroutine yang berjalan
		wg.Add(1)

		// 3. exec new goroutine
		go doPrint(&wg, data)
	}

	// 5. blocking
	// menunggu semua WaitGroup dieksekusi semuanya
	// jika ada 5 goroutine, maka ada 5 wg.done yang harus dipanggil
	wg.Wait()
}

// * Perbedaan WaitGroup dan Channel
// - channel tergantung pada goroutine tertentu, tidak seperti sync.WaitGroup yang tidak perlu tahu goroutine mana yang dijalankan (cukup tahu jumlah goroutine yang harus selesai)
// - penerapan sync.WaitGroup lebih mudah dibanding channel
// ---> channel untuk komunikasi data antar goroutine, sifatnya blocking bisa dimanfaatkan untuk manage goroutine;
// ---> WaitGroup khusus untuk sinkronisasi goroutine
// - perforam sync.WaitGroup lebih baik dibanding channel
