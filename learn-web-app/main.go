// * Hello World Web App and Static Routing

package main

import (
	"fmt"
	"net/http"
)

// method untuk handler route
// params w dan r
// - w sebagai mengoutputkan nilai return data, berbentuk []byte
// - r sebagai nilai input yang dibawa pada handler ini (query params, params, json)
func handlerIndex(w http.ResponseWriter, r *http.Request) {
	message := "welcome"
	w.Write([]byte(message))
}

func handlerHello(w http.ResponseWriter, r *http.Request) {
	message := "hello world!"
	w.Write([]byte(message))
}

func main() {
	// membuat route dengan HandleFunc
	// params 1 adalah routenya, dan params 2 adalah fungsi yang menghandle-nya
	// handler dapat berupa fungsi, closure, dan anonymous func
	// http.HandleFunc("/", handlerIndex)
	// http.HandleFunc("/index", handlerIndex)
	// http.HandleFunc("/hello", handlerHello)
	// nb: jika memanggil route yang tidak ada diatas, maka yang terpanggil adalah route root

	/* -------------------------------------------------------------------------- */
	// * static asset
	// akses ke static asset pake root bisa, namun akan mengacaukan route semuanya. karena root akan menjadi directory static asset
	// http.Handle("/", http.FileServer(http.Dir("assets")))

	// pake ini akan error, karena routing akan mengarah ke folder assets. yang menyebabkan route not found
	// http.Handle("/static", http.FileServer(http.Dir("assets")))

	// pake ini berhasil, karena dengan menambahkan StripPrefix(), akan menghapus prefix. sehingga setiap request pada route ini hanya akan di ambil url setelahnya
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("assets"))))

	address := ":6969"
	fmt.Println("server started at", address)

	// menjalankan server, bersifat blocking
	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
