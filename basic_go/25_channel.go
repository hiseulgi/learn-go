/* -------------------------------------------------------------------------- */
/*                                // * channel                                */
/* -------------------------------------------------------------------------- */
// channel digunakan untuk menghubungkan goroutine satu dengan yang lainnya
// jadi channel digunakan sebagai pengirim di satu goroutine
// dan satunya sebagai penerima di goroutine lainnya
// channel bersifat blocking / synchronous

// * keyword penting dalam channel adalah `make` dan `chan`
// var chan menjadi pengirim atau penerima dataa

package main

import "fmt"

func main() {
	// deklarasi channel string
	messages := make(chan string)

	// closure mengirim pesan string yang dikirim via channel
	sayHelloTo := func(who string) {
		data := fmt.Sprintf("hello %s", who)
		messages <- data // pross pengiriman data dari var data lewat channel messages
	}

	// ekseksusi closure pada goroutine berbeda (async)
	go sayHelloTo("andi")
	go sayHelloTo("sugab")
	go sayHelloTo("xenos")

	// ! ingat: eksekusi goroutine adalah async, sedangkan serah terima data antar channel adalah sync
	// fungsi yang diekseskusi paling awal akan diterima oleh var message1
	// proses penerimaan data dari channel messages ke var message1
	// proses ini bersifat blocking, artinya ia akan menunggu closure sayHelloTo dieksekusi terlebih dahulu

	// untuk mengambil data dari channel kita perlu tanda `<-`
	message1 := <-messages
	fmt.Println(message1)

	message2 := <-messages
	fmt.Println(message2)

	message3 := <-messages
	fmt.Println(message3)

	fmt.Println()

	/* -------------------------------------------------------------------------- */
	// * channel sebagai parameter
	// func funcName(paramName chan string)
	// jadi dengan fungsi yang berparameter channel, kita bisa menerima data channel tersebut di dalam fungsi
	for _, each := range []string{"ranggi ragata", "moona hoshi", "kobo kanaeru"} {
		// pengiriman data bersifat async
		go func(who string) {
			data := fmt.Sprintf("Woiii %s", who)
			messages <- data
		}(each)
	}

	// penerimanaan data bersifat sync (blocking)
	// dia bakal nunggu pesan terbaru
	for i := 0; i < 3; i++ {
		printMessage(messages)
	}
}

func printMessage(message chan string) {
	fmt.Println(<-message)
}
