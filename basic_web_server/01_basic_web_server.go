package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "ini index ya kak!")
}

func main() {
	// * basic web server
	// routing web menggunakan HandleFunc
	// 1 HandleFunc = 1 route
	// di dalam HandleFunc terdapat parameter yaitu nama route dan fungsi callbacknya (apa yang dilakukan didalamnya)
	// callback bisa iiie atau fungsi biasa
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello")
	})
	http.HandleFunc("/index", index)

	/* -------------------------------------------------------------------------- */
	// * template web
	// menggunakan file html sebagai template
	// dan kita bisa mengatur data-datanya
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]string{
			"Name":    "bagus",
			"Message": "have nice a day",
		}

		// parse file html apakah ada error penulisan
		t, err := template.ParseFiles("template.html")
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		// menerapkan data pada file html
		t.Execute(w, data)
	})

	// memulai menjalankan web server dengan ListenAndSerVe
	fmt.Println("starting web server at 127.0.0.1:8080")
	http.ListenAndServe(":8080", nil)
}
