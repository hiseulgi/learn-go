/* -------------------------------------------------------------------------- */
/*                           // * channel - timeout                           */
/* -------------------------------------------------------------------------- */
// mengontrol penerimaan data dari channel mengacu ke kapan waktu diterimanya data, dengan durasi timeout yang bisa ditentukan
// ketika tidak ada penerimaan data pada durasi tertentu, maka blok timeout akan dieksekusi

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// * send data to channel
func sendData(ch chan<- int) {
	// new randomizer obj
	randomizer := rand.New(rand.NewSource(time.Now().Unix()))

	// looping all
	for i := 0; true; i++ {
		// send data to channel
		ch <- i
		// randomize send data delay
		time.Sleep(time.Duration(randomizer.Int()%10+1) * time.Second)
	}
}

// * retrieve data from channel
func retreiveData(ch <-chan int) {
	// infinite retrieve data until no data after 5s
loop:
	for {
		select {
		// wait until new data from channel
		case data := <-ch:
			fmt.Println("receive data:", data)
		// if there no data after 5s, then timeout
		case <-time.After(time.Second * 5):
			fmt.Println("timeout. no activities under 5s")
			break loop
		}
	}
}

func main() {
	ch := make(chan int)

	// infinite send data by async
	go sendData(ch)

	// infinite retrieve data by sync (until no data in 5s)
	retreiveData(ch)

}
