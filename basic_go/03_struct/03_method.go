/* -------------------------------------------------------------------------- */
/*                                 // * method                                */
/* -------------------------------------------------------------------------- */
// method -> fungsi yang menempel pada type (struct atau tipe data lain)
// method bisa diakses lewat var objek
// dengan method kita memiliki akses ke property struct yang private

package main

import (
	"fmt"
	"strings"
)

// deklarasi type struct
type student struct {
	name  string
	grade int
}

// dekralasi method, print nama
func (s student) sayHello() {
	fmt.Println("halo", s.name)
}

// deklarasi method, mengambil kata ke berapa
func (s student) getNameAt(i int) string {
	return strings.Split(s.name, " ")[i-1]
}

// * method with pointer
// no pointer
// nilai tidak akan berubah, karena yang berubah hanya di dalam method ini aja
func (s student) changeName1(newName string) {
	fmt.Println(s.name, "has been changed into", newName)
	s.name = newName
}

// with pointer
// nilai akan berubah, karena mengacu pada alamat memori dari objek
func (s *student) changeName2(newName string) {
	fmt.Println(s.name, "has been changed into", newName)
	s.name = newName
}

func main() {
	s1 := student{
		"bagus adi prayoga",
		22,
	}

	s1.sayHello()
	nickname := s1.getNameAt(1)

	fmt.Println("panggilannya:", nickname)

	// ini bakal tetep nilainya
	s1.changeName1("asep udin") // bagus adi prayoga
	s1.sayHello()

	// ini bakal berubah nilainya
	s1.changeName2("andri anto") // andri anto
	s1.sayHello()
}
