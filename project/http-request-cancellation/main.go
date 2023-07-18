// HTTP Request Cancellation (Client HTTP Request)
// - kadangkala http request perlu waktu lama
// - di sisi client biasanya ada handler cancel request jika melebihi batas timeout

// pada program ini akan mencoba handle cancellation pada client http request dari sisi backend server

package main

import (
	"log"
	"net/http"
	"strings"
	"time"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	done := make(chan bool)
	go func() {
		// simulasi long-time request
		time.Sleep(10 * time.Second)
		done <- true
	}()

	select {
	// jika client membatalkan request
	case <-r.Context().Done():
		if err := r.Context().Err(); err != nil {
			if strings.Contains(strings.ToLower(err.Error()), "canceled") {
				log.Println("request canceled")
			} else {
				log.Println("unknown error occured.", err.Error())
			}
		}

	// jika server berhasil mengeksekusi program
	case <-done:
		log.Println("done")
	}
}

func main() {
	http.HandleFunc("/", IndexHandler)

	http.ListenAndServe(":8080", nil)
}
