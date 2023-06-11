// pada external file terdapat dua macam level akses
// * exported (public) dan unexported (private)
// level akses ini berlaku pada variabel, struct, method, dan property variabel

package library

import (
	"fmt"
)

// * variabel exported
var IsPass bool

// * struct
type Student struct {
	Name  string
	Grade int
	// grade int // unexported property
}

// * exported function (public function)
// ditandai dengan namanya yang menggunakan huruf kapital
// dapat digunakan diluar package ini
func SayHello(name string) {
	fmt.Println("hello")
	introduce(name) // kalau ini bisa dipanggil, soalnya masih dalem satu package
}

// * unexported function (private function)
// menggunakan huruf kecil
// hanya bisa dipakai di package ini saja
func introduce(name string) {
	fmt.Println("my name is", name)
}

// * fungsi init()
// fungsi yang otomatis di-run ketika di-import
func init() {
	fmt.Println("--> library/library.go imported bosqu!")
	IsPass = true
}
