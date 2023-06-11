/* -------------------------------------------------------------------------- */
/*                                 // * struct                                */
/* -------------------------------------------------------------------------- */
// berhubung di golang tidak memiliki class (OOP), maka hadirlah tipe data struktur (struct)

// struct -> kumpulan definisi variabel (property) dan/atau fungsi (method) yang dibukus sebagai tipe data baru
// variabel yang terbuat dari struct disebut object

package main

import "fmt"

type person struct {
	name string
	age  int
}

// * 11. nested struct
// buat type struct yang berisi struct juga
// cocok untuk decode data json yang kompleks
// * 12. tag proprerty dalam struct
// tag -> informasi opsional pada property struct, baisa dipake untuk encode/decode data json (tunggu bab nant)
// type teacher struct {
// 	person `tag1`
// 	info struct {
// 		address string `tag1`
// 		salary  int `tag1`
// 	}
// 	hobbies []string `tag1`
// }

// * 12 type alias
// struct bisa dibuatkan alias baru
type people = person

func main() {
	// * 8. kombinasi slice & struct
	persons := []person{
		{name: "nandi", age: 23},
		{name: "bagus", age: 20},
		{name: "adi", age: 21},
		{name: "wahyu", age: 21},
	}

	for _, person := range persons {
		fmt.Printf("%s berumur %d \n", person.name, person.age)
	}
	fmt.Println()

	// * 9. anonymous slice & struct
	students := []struct {
		person
		grade int
	}{
		{person: person{name: "asep", age: 20}, grade: 3},
		{person: person{name: "eva", age: 21}, grade: 2},
	}

	for _, student := range students {
		fmt.Println(student)
	}
	fmt.Println()

	// * 10. deklarasi anonymous struct dengan var
	var student struct {
		person
		grade int
	}
	student.person = person{"key", 22}
	student.age = 21
}
