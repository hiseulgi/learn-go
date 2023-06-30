/* -------------------------------------------------------------------------- */
/*                                 // * struct                                */
/* -------------------------------------------------------------------------- */
// berhubung di golang tidak memiliki class (OOP), maka hadirlah tipe data struktur (struct)

// struct -> kumpulan definisi variabel (property) dan/atau fungsi (method) yang dibukus sebagai tipe data baru
// variabel yang terbuat dari struct disebut object

package main

import "fmt"

// * 1. deklarasi struct
type student struct {
	name  string
	grade int
}

func main() {
	// * 2. penerapan struct
	var s1 student
	s1.name = "sugab"
	s1.grade = 3

	fmt.Println("student name:", s1.name)
	fmt.Println("student grade:", s1.grade)
	fmt.Println()

	// * 3. inisialisasi object struct
	// cara pertama kek diatas
	var s2 = student{"zia", 2}
	var s3 = student{name: "nandi"}
	// var s4 = student{name: "arip", grade: 3}
	// var s5 = student{grade: 2, name: "opet"}

	fmt.Println("s1:", s1.name)
	fmt.Println("s2:", s2.name)
	fmt.Println("s3:", s3.name)
	// fmt.Println("s4:", s4.name)
	// fmt.Println("s5:", s5.name)
	fmt.Println()

	// * 4. variabel objek pointer
	// membuat objek referensi dari objek s3
	var s4 *student = &s3

	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println("student3:", s3.name)
	fmt.Println("student4:", s4.name)

	s4.name = "laila"
	fmt.Println("student3:", s3.name)
	fmt.Println("student4:", s4.name)
}
