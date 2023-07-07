package main

import (
	// "access-level/library"
	// * import dengan prefix tanda titik
	// ini seperti pada python (from .. import *)
	// jadi ketika pemanggilan tidak perlu menulis nama packagenya terlebih dahulu
	. "access-level/library"

	// * alias pada import (as)
	// menyingkat nama package
	f "fmt"
)

func main() {
	student1 := Student{
		Name:  "sugab",
		Grade: 6,
		// grade: 4, // error unknown field
	}

	SayHello(student1.Name)
	f.Println("my grade is", student1.Grade)
	// library.introduce(name) // error undefined

	// mengakses func pada pacakge yang sama dari file partial.go
	// tapi ketika go build | go run harus menyertakan semua file dengan nama package
	// go run main.go partial.go
	// go run *.go

	// memanggil variabel dari library.go yang di-init
	if IsPass {
		SayGukGuk()
	}
}
