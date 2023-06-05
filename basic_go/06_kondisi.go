package main

import "fmt"

func main() {
	/* -------------------------------------------------------------------------- */
	/*                           // * 6. Seleksi Kondisi                          */
	/* -------------------------------------------------------------------------- */
	// * A. if else
	var grade = 7

	if grade == 10 {
		fmt.Println("Kamu jago bang!")
	} else if grade > 5 {
		fmt.Println("Kamu biasa aja!")
	} else {
		fmt.Println("Kamu cacat!")
	}

	// * B. Var Temp pada if else
	// var sementara pada blok kondisi, berguna agar lebih rapi dan var tsb bisa dipake pada kondisi2 dibawahnya
	var value = 86.0

	if percent := value / 100; percent >= 100 {
		fmt.Printf("%.1f jago!\n", percent)
	} else if percent >= 70 {
		fmt.Printf("%.1f biasa!\n", percent)
	} else {
		fmt.Printf("%.1f cacat!\n", percent)
	}

	// * C. switch case
	var point = 2

	switch point {
	case 8:
		fmt.Println("kamu jago")
	case 5, 6, 7:
		{
			fmt.Println("biasa aja lu")
			fmt.Println("hayuk")
		}
	default:
		fmt.Println("cacat")
	}

	// * D. switch dengan gaya if-else
	switch {
	case point == 8:
		fmt.Println("kamu jago")
	case (point < 8) && (point > 3):
		fmt.Println("biasa aja")
	default:
		{
			fmt.Println("cacat")
			fmt.Println("noob")
		}
	}

	// * E. fallthrough dalam switch
	// digunakan untuk memaksa mengecek kondisi case berikutnya dan langsung dieksekusi tanpa dicek nilai kondisinya
	point = 6

	switch {
	case point == 8:
		fmt.Println("kamu jago")
	case (point < 8) && (point > 3):
		fmt.Println("biasa aja")
		fallthrough
	default:
		{
			fmt.Println("cacat")
			fmt.Println("noob")
		}
	}
	// case kedua dan default akan tereksekusi
}
