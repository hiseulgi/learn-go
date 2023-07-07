package main

import "fmt"

func main() {
	/* -------------------------------------------------------------------------- */
	/*                                // * pointer                                */
	/* -------------------------------------------------------------------------- */
	// pointer -> reference atau alamt memori
	// var pointer -> var yang berisi alamat memori suatu nilai
	// niali default dari pointer adalah nil (kosong)

	// * 1. penerapan pointer
	// pembuatan pointer menggunakan tanda ( * )
	// var number *int
	// var name *string

	// referencing -> mengambil nilai pointer dari variabel biasa. caranya dengan menambah tanda ( & ) sebelum nama variabel
	// dereferencing -> mengambil nilai asli dari variabel pointer. caranya dengan menambah tanda ( * ) sebelum nama variabel

	var nameA string = "bagus"
	var nameB *string = &nameA

	fmt.Println("Nilai awal:")

	fmt.Println("nameA:", nameA)  // value
	fmt.Println("nameA:", &nameA) // address

	fmt.Println("nameB:", *nameB) // value
	fmt.Println("nameB:", nameB)  // address

	fmt.Println()

	// * 2. efek perubahan pointer
	*nameB = "sugab"

	fmt.Println("Nilai setelah diubah:")

	fmt.Println("nameA:", nameA)  // value
	fmt.Println("nameA:", &nameA) // address

	fmt.Println("nameB:", *nameB) // value
	fmt.Println("nameB:", nameB)  // address

	fmt.Println()

	// * 3. parameter pointer
	number := 10

	fmt.Println("number awal:", number)

	change(&number, 69)

	fmt.Println("number setelah change:", number)
	fmt.Println()
}

func change(original *int, value int) {
	*original = value
}
