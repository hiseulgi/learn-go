package main

import "fmt"

func main() {
	/* -------------------------------------------------------------------------- */
	/*                             // * 7. Perulangan                             */
	/* -------------------------------------------------------------------------- */
	// * for (di go hanya ada for aja, cuma dia bisa dipake buat for, foreach, while)
	// biasa
	for i := 0; i < 5; i++ {
		fmt.Println("angka", i)
	}
	fmt.Println()

	// only condition
	var i = 0
	for i < 5 {
		fmt.Println("angka", i)
		i++
	}
	fmt.Println()

	// no argument
	i = 0
	for {
		fmt.Println("angka", i)
		i++
		if i == 5 {
			break
		}
	}
	fmt.Println()

	// break & continue
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}

		if i > 7 {
			break
		}
		fmt.Println("angka", i)
	}
	fmt.Println()

	// for bersarang
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
	fmt.Println()

	// label in for
	// dengan kita bisa melakukan break dan continue pada blok for terluar
outerloop:
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			if i == 3 {
				break outerloop
			}
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
	fmt.Println()
}
