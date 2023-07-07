// * MUTEX
// prolog:
// race condition -> kondisi dimana lebih dari satu goroutine mengakses data yang sama pada waktu yang tepat bersamaan. ketika hal ini terjadi, nilai data tersebut akan menjadi kacau (sering terjadi pada concurrency programming)
// mutex melakukan pengubahan level akses sebuah data menjadi ekslusif, menjadi data tersebut hanya dapat dikonsumsi (r/w) oleh satu buah goroutine saja.
// jadi hanya 1 goroutine saja yang dapat mengakses data tersebut, dan goroutine lain akan dipaksa menunggu

// dengan sync.Mutex kita dapat melakukan lock dan unlock data, sehingga race condition dapat teratasi

// * cara deteksi race condition
// pada saat run tambahi dengan -race
// go run -race *.go

package main

import (
	"fmt"
	"sync"
)

type counter struct {
	// perlu menambah embed struct sync.Mutex
	sync.Mutex
	val int
}

// untuk menggunakan mutex, disini kita tambahkan method Lock dan Unlock
// coba jalankan tanpa c.Lock() dan c.Unlock() pasti terjadi race condition
func (c *counter) add() {
	// c.Lock() // baris ini ekslusif, hanya 1 goroutine saja yang dapat mengakses nilai c dibawah
	c.val++
	// c.Unlock() // membuka kembali akses pada var yang di lock
}

func (c *counter) get() int {
	return c.val
}

func main() {
	wg := sync.WaitGroup{}
	meter := counter{}

	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func() {
			for j := 0; j < 100; j++ {
				meter.add()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(meter.get())
}
