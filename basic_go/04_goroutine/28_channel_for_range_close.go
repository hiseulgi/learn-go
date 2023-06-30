/* -------------------------------------------------------------------------- */
/*                       // * channel - range and close                       */
/* -------------------------------------------------------------------------- */
// proses retrieving data dari banyak channel lebih mudah jika menggunakan for - range
// for - range berfungsi untuk handle penerimaan data
// setiap ada pengiriman data via channel, akan men-trigger for - range yang akan berjalan terus menerus
// perulangan hanya akan berhenti jika channel di-`close`
// channel yang di-`close` tidak bisa digunakan lagi untuk menerima ataupun mengirim data

package main

import "fmt"

// func to send data into channel
// param chan untuk mengirim data
func sendData(word string, iterate int, channel chan<- string) {
	// send data to channel with for loop
	for i := 0; i < iterate; i++ {
		wordToSend := fmt.Sprintf("%s-%d", word, i)
		channel <- wordToSend // blocking process
	}

	// after done, channel will be closed
	close(channel)
}

// func to receive data from channel
// param chan untuk menerima data
func receiveData(channel <-chan string) {
	for message := range channel {
		fmt.Println(message)
	}
}

func main() {
	// make string channel variable
	messages := make(chan string)
	iterate := 20

	// send data on channel messages by async
	go sendData("anjay", iterate, messages)

	// retrieve data from channel messages by sync
	receiveData(messages)
}

// * channel direction
// tujuannya untuk menentukan level akses channel pada parameter
/* -------------------------------------------------------------------------- */
// `ch chan string`		: parameter mengirim dan menerima data
// `ch chan<- string`	: parameter untuk mengirim data
// `ch <-chan string`	: parameter untuk menerima data
/* 	-------------------------------------------------------------------------- */
