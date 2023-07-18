// cookie -> data dalam bentuk teks yang disimpan pada browser ketika pengunjung sedang surfing ke situs

package main

import (
	"fmt"
	"net/http"
	"time"

	gubrak "github.com/novalagung/gubrak/v2"
)

type M map[string]interface{}

var cookieName = "CookieData"

func main() {
	http.HandleFunc("/", ActionIndex)
	http.HandleFunc("/delete", ActionDelete)

	fmt.Println("run server at :8080")
	http.ListenAndServe(":8080", nil)
}

func ActionIndex(w http.ResponseWriter, r *http.Request) {
	c := &http.Cookie{}

	// mengakses cookie sekarang
	if storedCookie, _ := r.Cookie(cookieName); storedCookie != nil {
		c = storedCookie
	}

	// jika cookie kosong, maka buat baru
	if c.Value == "" {
		c = &http.Cookie{}
		c.Name = cookieName
		c.Value = gubrak.RandomString(32)
		c.Expires = time.Now().Add(5 * time.Minute)
		http.SetCookie(w, c)
	}

	w.Write([]byte(c.Value))
}

func ActionDelete(w http.ResponseWriter, r *http.Request) {
	c := &http.Cookie{}
	c.Name = cookieName
	c.Expires = time.Unix(0, 0)
	c.MaxAge = -1
	http.SetCookie(w, c)

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
