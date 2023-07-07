/* -------------------------------------------------------------------------- */
/*                                 // * struct                                */
/* -------------------------------------------------------------------------- */
// berhubung di golang tidak memiliki class (OOP), maka hadirlah tipe data struktur (struct)

// struct -> kumpulan definisi variabel (property) dan/atau fungsi (method) yang dibukus sebagai tipe data baru
// variabel yang terbuat dari struct disebut object

package main

import "fmt"

// * 5. embedded struct (struct di dalam struct)
type person struct {
	name string
	age  int
}

type student struct {
	person
	grade int
	// age   int
}

func main() {
	var s1 = student{}
	s1.name = "sugab"
	s1.age = 20
	// s1.person.age = 22
	s1.grade = 3

	fmt.Println(s1)
	fmt.Println(s1.name)
	fmt.Println(s1.age)
	fmt.Println(s1.person.age)
	fmt.Println(s1.grade)
	fmt.Println()

	// * 6. pengisian nilai sub-struct
	// cara lain mengisi sub-structnya
	var p2 = person{name: "arip", age: 25}
	var s2 = student{person: p2, grade: 2}

	fmt.Println(s2)
	fmt.Println()

	// * 7. anonymous struct
	// cocok untuk sekali pakai
	var dosen1 = struct {
		person
		class string
	}{
		// mengisi property langsung
		// person: person{"edi", 30},
		// class:  "math",
	}
	dosen1.person = person{"edi", 30}
	dosen1.class = "math"

	fmt.Println(dosen1)
	fmt.Println()
}
